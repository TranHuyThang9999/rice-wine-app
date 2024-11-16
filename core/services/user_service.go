package services

import (
	"context"
	"rice-wine-shop/common/log"
	"rice-wine-shop/common/utils"
	"rice-wine-shop/core/apperrors"
	"rice-wine-shop/core/domain"
	"rice-wine-shop/core/entities"
	"rice-wine-shop/core/enums"

	"gorm.io/gorm"
)

type UserService struct {
	user  domain.RepositoryUser
	file  domain.RepositoryFileStore
	trans domain.TransactionHelper
}

func NewServiceUser(user domain.RepositoryUser,
	file domain.RepositoryFileStore,
	trans domain.TransactionHelper,
) *UserService {
	return &UserService{
		user:  user,
		file:  file,
		trans: trans,
	}
}

func (u *UserService) AddUser(ctx context.Context, req *entities.CreateUsersRequest) error {
	userID := utils.GenerateUniqueKey()

	err := u.trans.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		hashPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return err
		}
		err = u.user.Create(ctx, tx, &domain.User{
			ID:          userID,
			Name:        req.Name,
			PhoneNumber: req.PhoneNumber,
			Password:    hashPassword,
			Role:        enums.ROLE_USER,
			CreatedAt:   utils.GenerateTimestamp(),
			UpdatedAt:   utils.GenerateTimestamp(),
		})
		if err != nil {
			log.Error(err, "error")
			return err
		}
		err = u.file.Create(ctx, tx, &domain.FileStore{
			ID:    utils.GenerateUniqueKey(),
			AnyID: userID,
			Path:  req.Avatar,
		})
		if err != nil {
			log.Error(err, "error")
			return err
		}

		return nil
	})

	if err != nil {
		log.Error(err, "error")
		return err
	}
	return nil
}
func (u *UserService) ProFileUser(ctx context.Context, phoneNumber string) (*entities.User, error) {
	user, err := u.user.GetUserByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		log.Error(err, "error")
		return nil, err
	}
	if user == nil {
		return nil, apperrors.ErrUserNotFound
	}
	files, err := u.file.GetListFileByObjectID(ctx, user.ID)
	if err != nil {
		log.Error(err, "error")
		return nil, err
	}
	pathAvatar := ""
	if len(files) == 0 {
		pathAvatar = ""
	} else {
		pathAvatar = files[0].Path
	}
	return &entities.User{
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Avatar:      pathAvatar,
		Role:        user.Role,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}
