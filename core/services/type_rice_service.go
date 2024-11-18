package services

import (
	"context"
	"fmt"
	"rice-wine-shop/common/log"
	"rice-wine-shop/common/utils"
	"rice-wine-shop/core/apperrors"
	"rice-wine-shop/core/domain"
	"rice-wine-shop/core/entities"
	"strings"

	"gorm.io/gorm"
)

type TypeRiceService struct {
	typeRice domain.RepositoryTypeRice
	file     domain.RepositoryFileStore
	trans    domain.TransactionHelper
}

func NewTypeRiceService(typeRice domain.RepositoryTypeRice,
	trans domain.TransactionHelper,
	file domain.RepositoryFileStore) *TypeRiceService {
	return &TypeRiceService{
		typeRice: typeRice,
		file:     file,
		trans:    trans,
	}
}

func (u *TypeRiceService) AddTypeRice(ctx context.Context, creatorID int64, req *entities.CreateTypeRiceRequest) (*apperrors.ErrTypeRice, error) {
	if u.trans == nil {
		return nil, fmt.Errorf("transaction manager is not initialized")
	}
	count, err := u.typeRice.GetTypeRiceNameByUserID(ctx, creatorID, strings.TrimSpace(req.Name))
	if err != nil {
		log.Error(err, "error")
		return nil, err
	}
	if count > 0 {
		return apperrors.ErrConflictTypeName.Pointer(), nil
	}
	err = u.trans.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		typeRiceID := utils.GenerateUniqueKey()
		err := u.typeRice.Add(ctx, tx, &domain.TypeRice{
			CreatorID: creatorID,
			ID:        typeRiceID,
			Name:      strings.TrimSpace(req.Name),
		})
		if err != nil {
			log.Error(err, "error")
			return err
		}
		if len(req.Files) > 0 {
			var listFile = make([]*domain.FileStore, 0)
			for _, file := range req.Files {
				listFile = append(listFile, &domain.FileStore{
					ID:    utils.GenerateUniqueKey(),
					AnyID: typeRiceID,
					Path:  file,
				})
			}
			err = u.file.AddListFile(ctx, tx, listFile)
			if err != nil {
				log.Error(err, "error")
				return err
			}
		}

		return nil
	})
	if err != nil {
		log.Error(err, "error")
		return nil, err
	}
	return nil, err
}

func (u *TypeRiceService) GetTypeRiceNameByUserID(ctx context.Context, userID int64) ([]*entities.ListTypeRiceResponse, error) {
	typeRiceList, err := u.typeRice.GetListByCreator(ctx, userID)
	if err != nil {
		log.Error(err, "failed to fetch list of type rice")
		return nil, err
	}

	allFiles, err := u.file.GetListFileByUserID(ctx, userID)
	if err != nil {
		log.Error(err, "failed to fetch files for type rice")
		return nil, err
	}

	fileMap := make(map[int64][]*domain.FileStore)
	for _, file := range allFiles {
		fileMap[file.AnyID] = append(fileMap[file.AnyID], file)
	}

	responses := make([]*entities.ListTypeRiceResponse, 0, len(typeRiceList))
	for _, typeRice := range typeRiceList {
		responses = append(responses, &entities.ListTypeRiceResponse{
			ID:    typeRice.ID,
			Name:  typeRice.Name,
			Files: fileMap[typeRice.ID],
		})
	}

	return responses, nil
}

func (u *TypeRiceService) DeleteById(ctx context.Context, id int64) error {
	if err := u.trans.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		err := u.typeRice.DeleteById(ctx, tx, id)
		if err != nil {
			log.Error(err, "error")
			return err
		}
		err = u.file.DeleteListFileByObjectID(ctx, tx, id)
		if err != nil {
			log.Error(err, "error")
			return err
		}
		return nil
	}); err != nil {
		log.Error(err, "error trans")
		return err
	}

	return nil
}

func (u *TypeRiceService) UpdateById(ctx context.Context, userID int64, req *entities.UpdateTypeRiceRequest) (*apperrors.ErrTypeRice, error) {
	existingTypeRice, err := u.typeRice.GetByName(ctx, userID, req.Name)
	if err != nil {
		log.Error(err, "error")
		return nil, err
	}
	if existingTypeRice != nil && existingTypeRice.ID != req.ID {
		log.Info("type rice name already exists")
		return apperrors.ErrConflictTypeName.Pointer(), nil
	}
	err = u.typeRice.UpdateById(ctx, &domain.TypeRice{
		ID:        req.ID,
		Name:      req.Name,
		CreatorID: userID,
	})
	if err != nil {
		log.Error(err, "error")
		return nil, err
	}

	return nil, nil
}
