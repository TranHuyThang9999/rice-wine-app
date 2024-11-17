package repository

import (
	"context"
	"rice-wine-shop/core/domain"

	"gorm.io/gorm"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) domain.RepositoryFileStore {
	return &FileRepository{
		db: db,
	}
}

func (f *FileRepository) AddListFile(ctx context.Context, tx *gorm.DB, req []*domain.FileStore) error {
	result := tx.WithContext(ctx).Create(&req)
	return result.Error
}

func (f *FileRepository) Create(ctx context.Context, tx *gorm.DB, req *domain.FileStore) error {
	result := tx.WithContext(ctx).Create(&req)
	return result.Error
}

func (f *FileRepository) DeleteById(ctx context.Context, id int64) error {
	panic("unimplemented")
}

func (f *FileRepository) GetListFileByObjectID(ctx context.Context, anyID int64) ([]*domain.FileStore, error) {
	var files = make([]*domain.FileStore, 0)
	result := f.db.WithContext(ctx).Where("any_id = ?", anyID).Find(&files)
	return files, result.Error
}
func (r *FileRepository) GetListFileByUserID(ctx context.Context, userID int64) ([]*domain.FileStore, error) {
	var files []*domain.FileStore
	err := r.db.WithContext(ctx).
		Table("file_stores").
		Select("file_stores.id, file_stores.any_id, file_stores.path").
		Joins("JOIN type_rices ON file_stores.any_id = type_rices.id").
		Where("type_rices.creator_id = ?", userID).
		Find(&files).Error
	if err != nil {
		return nil, err
	}
	return files, nil
}
