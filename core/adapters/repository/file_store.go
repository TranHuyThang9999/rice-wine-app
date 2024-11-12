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
	result := tx.Create(&req)
	return result.Error
}

func (f *FileRepository) Create(ctx context.Context, tx *gorm.DB, req *domain.FileStore) error {
	result := tx.Create(&req)
	return result.Error
}

func (f *FileRepository) DeleteById(ctx context.Context, id int64) error {
	panic("unimplemented")
}

func (f *FileRepository) GetListFileByObjectID(ctx context.Context, anyID int64) ([]*domain.FileStore, error) {
	var files = make([]*domain.FileStore, 0)
	result := f.db.Where("any_id = ?", anyID).Find(&files)
	return files, result.Error
}
