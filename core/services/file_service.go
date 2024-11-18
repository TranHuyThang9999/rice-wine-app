package services

import (
	"context"
	"rice-wine-shop/common/log"
	"rice-wine-shop/common/utils"
	"rice-wine-shop/core/apperrors"
	"rice-wine-shop/core/domain"
	"rice-wine-shop/core/entities"
)

type FileStoreSerVice struct {
	file domain.RepositoryFileStore
}

func NewFileStoreSerVice(file domain.RepositoryFileStore) *FileStoreSerVice {
	return &FileStoreSerVice{
		file: file,
	}
}

func (u *FileStoreSerVice) DeleteFileByID(ctx context.Context, fileID, userID int64) error {
	err := u.file.DeleteById(ctx, fileID, userID)
	if err != nil {
		log.Error(err, "error")
		return err
	}
	return nil
}

func (u *FileStoreSerVice) AddListFile(ctx context.Context, userID int64, req *entities.CreateUploadFileRequest) error {
	listFileUpload := make([]*domain.FileStore, 0)
	if len(req.Paths) == 0 {
		return apperrors.ErrorUploadFiles
	}

	for _, v := range req.Paths {
		listFileUpload = append(listFileUpload, &domain.FileStore{
			ID:        utils.GenerateUniqueKey(),
			AnyID:     req.ObjectID,
			Path:      v,
			CreatorID: userID,
		})
	}
	err := u.file.UpsetFiles(ctx, listFileUpload)
	if err != nil {
		log.Error(err, "error")
		return err
	}

	return nil

}
