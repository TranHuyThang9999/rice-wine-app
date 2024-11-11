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

func (f *FileRepository) AddListFile(ctx context.Context, req []*domain.FileStore) error {
	panic("unimplemented")
}

func (f *FileRepository) Create(ctx context.Context, tx *gorm.DB, req *domain.FileStore) error {
	result := tx.Create(&req)
	return result.Error
}

func (f *FileRepository) DeleteById(ctx context.Context, id int64) error {
	panic("unimplemented")
}

func (f *FileRepository) GetListFileByObjectID(ctx context.Context, anyID int64) ([]*domain.FileStore, error) {
	panic("unimplemented")
}
