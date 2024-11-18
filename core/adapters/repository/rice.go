package repository

import (
	"context"
	"rice-wine-shop/core/domain"

	"gorm.io/gorm"
)

type RiceRepository struct {
	db *gorm.DB
}

func NewRiceRepository(db *gorm.DB) domain.RepositoryRice {
	return &RiceRepository{db: db}
}

func (r *RiceRepository) Create(ctx context.Context, tx *gorm.DB, req *domain.Rices) error {
	if err := tx.WithContext(ctx).Create(req).Error; err != nil {
		return err
	}
	return nil
}

func (r *RiceRepository) GetListByCreatorID(ctx context.Context, creatorID int64) ([]*domain.Rices, error) {
	var riceList []*domain.Rices
	if err := r.db.WithContext(ctx).
		Where("creator_id = ?", creatorID).
		Find(&riceList).Error; err != nil {
		return nil, err
	}
	return riceList, nil
}

func (r *RiceRepository) DeleteById(ctx context.Context, id int64) error {
	if err := r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&domain.Rices{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *RiceRepository) UpdateById(ctx context.Context, req *domain.Rices) error {
	if err := r.db.WithContext(ctx).
		Where("id = ?", req.ID).
		Updates(req).Error; err != nil {
		return err
	}
	return nil
}

func (r *RiceRepository) GetByRiceName(ctx context.Context, userID int64, name string) (int64, error) {
	var count int64
	result := r.db.WithContext(ctx).Model(&domain.Rices{}).Where("creator_id = ? and name = ?", userID, name).Count(&count)
	return count, result.Error
}

func (r *RiceRepository) GetListFileByUserID(ctx context.Context, userID int64) ([]*domain.FileStore, error) {
	var files []*domain.FileStore
	err := r.db.WithContext(ctx).
		Table("file_stores").
		Select("file_stores.id, file_stores.any_id, file_stores.path").
		Joins("JOIN rices ON file_stores.any_id = rices.id").
		Where("rices.creator_id = ?", userID).
		Find(&files).Error
	if err != nil {
		return nil, err
	}
	return files, nil
}
