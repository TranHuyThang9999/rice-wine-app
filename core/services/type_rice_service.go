package services

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"rice-wine-shop/common/log"
	"rice-wine-shop/common/utils"
	"rice-wine-shop/core/domain"
	"rice-wine-shop/core/entities"
	"strings"
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

func (u *TypeRiceService) AddTypeRice(ctx context.Context, creatorID int64, req *entities.CreateTypeRiceRequest) error {
	if u.trans == nil {
		return fmt.Errorf("transaction manager is not initialized")
	}
	count, err := u.typeRice.GetTypeRiceNameByUserID(ctx, creatorID, strings.TrimSpace(req.Name))
	if err != nil {
		log.Error(err, "error")
		return err
	}
	if count > 0 {
		return fmt.Errorf("type rice name %s already exists", req.Name)
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
		return nil
	})
	if err != nil {
		log.Error(err, "error")
		return err
	}
	return nil
}
