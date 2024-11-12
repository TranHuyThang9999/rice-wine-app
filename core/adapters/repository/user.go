package repository

import (
	"context"
	"errors"
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
	result := tx.WithContext(ctx).Create(&req)
	return result.Error
}

func (u *UserRepository) GetLister(ctx context.Context) ([]*domain.User, error) {
	var users = make([]*domain.User, 0)
	result := u.db.WithContext(ctx).Find(&users)
	return users, result.Error
}

func (u *UserRepository) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*domain.User, error) {
	var user *domain.User
	result := u.db.WithContext(ctx).Where("phone_number = ?", phoneNumber).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepository) UpdateUserById(ctx context.Context, tx *gorm.DB, req *domain.User) error {
	result := tx.WithContext(ctx).Save(req)
	return result.Error
}
