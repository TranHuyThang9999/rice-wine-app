package services

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"rice-wine-shop/common/log"
	"rice-wine-shop/common/utils"
	"rice-wine-shop/core/apperrors"
	"rice-wine-shop/core/domain"
	"rice-wine-shop/core/entities"
)

type RiceService struct {
	rice     domain.RepositoryRice
	file     domain.RepositoryFileStore
	tran     domain.TransactionHelper
	typeRice domain.RepositoryTypeRice
}

func NewRiceService(rice domain.RepositoryRice,
	file domain.RepositoryFileStore,
	tran domain.TransactionHelper,
	typeRice domain.RepositoryTypeRice) *RiceService {
	return &RiceService{
		rice:     rice,
		file:     file,
		tran:     tran,
		typeRice: typeRice,
	}
}

func (u *RiceService) AddRice(ctx context.Context, userID int64, req *entities.CreateRiceRequest) (*apperrors.ErrTypeRice, error) {
	if u.tran == nil {
		return nil, fmt.Errorf("transaction manager is not initialized")
	}
	countRiceName, err := u.rice.GetByRiceName(ctx, userID, req.Name)
	if err != nil {
		return nil, err
	}
	if countRiceName > 0 {
		return apperrors.ErrConflictTypeName.Pointer(), nil
	}
	riceID := utils.GenerateUniqueKey()
	err = u.tran.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		count, err := u.typeRice.CheckExistsTypeRiceByID(ctx, userID, req.TypeRiceID)
		if err != nil {
			return err
		}
		if count == 0 {
			return fmt.Errorf("type rice with id %v not found", req.TypeRiceID)
		}
		err = u.rice.Create(ctx, tx, &domain.Rices{
			ID:            riceID,
			CreatorID:     userID,
			TypeRiceID:    req.TypeRiceID,
			Name:          req.Name,
			Quantity:      req.Quantity,
			Price:         req.Price,
			Origin:        req.Origin,
			HarvestSeason: req.HarvestSeason,
			CreatedAt:     utils.GenerateTimestamp(),
			UpdatedAt:     utils.GenerateTimestamp(),
		})
		if err != nil {
			log.Error(err, "error creating rice")
			return err
		}

		if len(req.Files) > 0 {
			var listFile = make([]*domain.FileStore, 0)
			for _, file := range req.Files {
				listFile = append(listFile, &domain.FileStore{
					ID:    utils.GenerateUniqueKey(),
					AnyID: riceID,
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
		return nil, err
	}
	return nil, nil
}

func (u *RiceService) GetListRiceByUserID(ctx context.Context, userID int64) ([]*entities.ListRiceByUserIDResponse, error) {
	listRiceByUserID, err := u.rice.GetListByCreatorID(ctx, userID)
	if err != nil {
		return nil, err
	}

	allFiles, err := u.rice.GetListFileByUserID(ctx, userID)
	if err != nil {
		log.Error(err, "failed to fetch files for type rice")
		return nil, err
	}
	fileMap := make(map[int64][]*domain.FileStore)
	for _, file := range allFiles {
		fileMap[file.AnyID] = append(fileMap[file.AnyID], file)
	}
	responses := make([]*entities.ListRiceByUserIDResponse, 0, len(listRiceByUserID))
	for _, typeRice := range listRiceByUserID {
		responses = append(responses, &entities.ListRiceByUserIDResponse{
			ID:    typeRice.ID,
			Name:  typeRice.Name,
			Files: fileMap[typeRice.ID],
		})
	}
	return responses, nil
}
