package repository

import (
	"context"
	"rice-wine-shop/core/domain"

	"gorm.io/gorm"
)

type TypeRiceRepository struct {
	db *gorm.DB
}

func NewTypeRiceRepository(db *gorm.DB) domain.RepositoryTypeRice {
	return &TypeRiceRepository{db: db}
}

func (t *TypeRiceRepository) GetListByCreator(ctx context.Context, creatorID int64) ([]*domain.TypeRice, error) {
	var list = make([]*domain.TypeRice, 0)
	result := t.db.WithContext(ctx).Where("creator_id = ?", creatorID).Find(&list)
	return list, result.Error
}

func (t *TypeRiceRepository) Add(ctx context.Context, tx *gorm.DB, req *domain.TypeRice) error {
	result := tx.WithContext(ctx).Create(req)
	return result.Error
}

func (t *TypeRiceRepository) DeleteById(ctx context.Context, tx *gorm.DB, id int64) error {
	result := tx.WithContext(ctx).Where("id = ?", id).Delete(&domain.TypeRice{})
	return result.Error
}

func (t *TypeRiceRepository) UpdateById(ctx context.Context, req *domain.TypeRice) error {
	result := t.db.WithContext(ctx).Save(&req)
	return result.Error
}

func (t *TypeRiceRepository) GetTypeRiceNameByUserID(ctx context.Context, userID int64, nameRice string) (int64, error) {
	var count int64
	result := t.db.WithContext(ctx).Model(&domain.TypeRice{}).Where("creator_id = ? and name = ?", userID, nameRice).Count(&count)
	return count, result.Error
}

func (t *TypeRiceRepository) CheckExistsTypeRiceByID(ctx context.Context, userID int64, typeRiceID int64) (int64, error) {
	var count int64
	result := t.db.WithContext(ctx).Model(&domain.TypeRice{}).Where("creator_id = ? and id = ?", userID, typeRiceID).Count(&count)
	return count, result.Error
}

func (t *TypeRiceRepository) GetByName(ctx context.Context, userID int64, nameRice string) (*domain.TypeRice, error) {
	var typeRice *domain.TypeRice
	result := t.db.WithContext(ctx).Where("creator_id = ? and name = ?", userID, nameRice).First(&typeRice)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return typeRice, result.Error
}
