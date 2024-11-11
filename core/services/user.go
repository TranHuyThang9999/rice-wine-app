package services

import (
	"context"
	"rice-wine-shop/common/log"
	"rice-wine-shop/common/utils"
	"rice-wine-shop/core/constant"
	"rice-wine-shop/core/domain"
	"rice-wine-shop/core/entities"

	"gorm.io/gorm"
)

type ServiceUser struct {
	user  domain.RepositoryUser
	file  domain.RepositoryFileStore
	trans domain.TransactionHelper
}

func NewServiceUser(user domain.RepositoryUser,
	file domain.RepositoryFileStore,
	trans domain.TransactionHelper,
) *ServiceUser {
	return &ServiceUser{
		user:  user,
		file:  file,
		trans: trans,
	}
}

func (u *ServiceUser) AddUser(ctx context.Context, req *entities.CreateUsersRequest) error {
	userID := utils.GenerateUniqueKey()
	err := u.trans.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		err := u.user.Create(ctx, tx, &domain.User{
			ID:          userID,
			Name:        req.Name,
			PhoneNumber: req.PhoneNumber,
			Password:    req.Password,
			Role:        constant.ROLE_USER,
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
