package repository

import (
	"context"
	"gorm.io/gorm"
	"rice-wine-shop/core/domain"
)

type TypeRiceRepository struct {
	db *gorm.DB
}

func NewTypeRiceRepository(db *gorm.DB) domain.RepositoryTypeRice {
	return &TypeRiceRepository{db: db}
}

func (t *TypeRiceRepository) Add(ctx context.Context, tx *gorm.DB, req *domain.TypeRice) error {
	result := tx.WithContext(ctx).Create(req)
	return result.Error
}

func (t *TypeRiceRepository) DeleteById(ctx context.Context, id int64) error {
	result := t.db.WithContext(ctx).Where("id = ?", id).Delete(&domain.TypeRice{})
	return result.Error
}

func (t *TypeRiceRepository) UpdateById(ctx context.Context, req *domain.TypeRice) error {
	result := t.db.WithContext(ctx).Save(req)
	return result.Error
}

func (t *TypeRiceRepository) GetList(ctx context.Context) ([]*domain.TypeRice, error) {
	var list = make([]*domain.TypeRice, 0)
	result := t.db.WithContext(ctx).Find(&list)
	return list, result.Error
}
func (t *TypeRiceRepository) GetTypeRiceNameByUserID(ctx context.Context, userID int64, nameRice string) (int64, error) {
	var count int64
	result := t.db.WithContext(ctx).Model(&domain.TypeRice{}).Where("creator_id = ? and name = ?", userID, nameRice).Count(&count)
	return count, result.Error
}
