package domain

import (
	"context"
	"gorm.io/gorm"
)

type TypeRice struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	CreatorID int64          `json:"creatorId"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type RepositoryTypeRice interface {
	Add(ctx context.Context, tx *gorm.DB, req *TypeRice) error
	DeleteById(ctx context.Context, id int64) error
	UpdateById(ctx context.Context, req *TypeRice) error
	GetList(ctx context.Context) ([]*TypeRice, error)
	GetTypeRiceNameByUserID(ctx context.Context, userID int64, nameRice string) (int64, error)
}
