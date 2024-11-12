package domain

import (
	"context"

	"gorm.io/gorm"
)

type FileStore struct {
	ID    int64  `json:"id,omitempty"`
	AnyID int64  `json:"anyId,omitempty"`
	Path  string `json:"path,omitempty"`
}
type RepositoryFileStore interface {
	Create(ctx context.Context, tx *gorm.DB, req *FileStore) error
	AddListFile(ctx context.Context, tx *gorm.DB, req []*FileStore) error
	GetListFileByObjectID(ctx context.Context, anyID int64) ([]*FileStore, error)
	DeleteById(ctx context.Context, id int64) error
}
