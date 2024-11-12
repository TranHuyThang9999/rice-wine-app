package domain

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Password    string `json:"password"`
	Email       string `json:"email,omitempty"`
	Role        int    `json:"role,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	UpdatedAt   int64  `json:"updated_at,omitempty"`
}

type RepositoryUser interface {
	Create(ctx context.Context, tx *gorm.DB, req *User) error
	GetLister(ctx context.Context) ([]*User, error)
	GetUserByPhoneNumber(ctx context.Context, name string) (*User, error)
	UpdateUserById(ctx context.Context, tx *gorm.DB, req *User) error
}
