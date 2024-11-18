package domain

import (
	"context"

	"gorm.io/gorm"
)

type TypeRice struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	CreatorID int64          `json:"creator_id"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type RepositoryTypeRice interface {
	Add(ctx context.Context, tx *gorm.DB, req *TypeRice) error
	DeleteById(ctx context.Context, tx *gorm.DB, id int64) error
	UpdateById(ctx context.Context, req *TypeRice) error
	GetListByCreator(ctx context.Context, creatorID int64) ([]*TypeRice, error)
	GetTypeRiceNameByUserID(ctx context.Context, userID int64, nameRice string) (int64, error)
	CheckExistsTypeRiceByID(ctx context.Context, userID int64, typeRiceID int64) (int64, error)
}
