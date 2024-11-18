package domain

import (
	"context"

	"gorm.io/gorm"
)

type FileStore struct {
	ID        int64  `json:"id,omitempty"`
	AnyID     int64  `json:"anyId,omitempty"`
	Path      string `json:"path,omitempty"`
	CreatorID int64  `json:"creatorId,omitempty"`
}
type RepositoryFileStore interface {
	Create(ctx context.Context, tx *gorm.DB, req *FileStore) error
	AddListFile(ctx context.Context, tx *gorm.DB, req []*FileStore) error
	GetListFileByObjectID(ctx context.Context, anyID int64) ([]*FileStore, error) //objectID khác userID
	DeleteById(ctx context.Context, fileId, userID int64) error
	GetListFileByUserID(ctx context.Context, userID int64) ([]*FileStore, error)
	DeleteListFileByObjectID(ctx context.Context, tx *gorm.DB, id int64) error
	UpsetFiles(ctx context.Context, req []*FileStore) error
}
