package repository

import (
	"context"
	"rice-wine-shop/core/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.RepositoryUser {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(ctx context.Context, tx *gorm.DB, req *domain.User) error {
	result := tx.Create(&req)
	return result.Error
}

func (u *UserRepository) GetListuser(ctx context.Context) ([]*domain.User, error) {
	var users = make([]*domain.User, 0)
	result := u.db.Find(&users)
	return users, result.Error
}

func (u *UserRepository) GetuserByPhoneNumber(ctx context.Context, phone_number string) (*domain.User, error) {
	var user *domain.User
	result := u.db.WithContext(ctx).Where("phone_number = ?", phone_number).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepository) UpdateUserById(ctx context.Context, req *domain.User) error {
	panic("unimplemented")
}
