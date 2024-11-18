package entities

import (
	"rice-wine-shop/core/domain"
)

type CreateTypeRiceRequest struct {
	Name  string   `json:"name" binding:"required"`
	Files []string `json:"files"`
}

type ListTypeRiceResponse struct {
	ID    int64               `json:"id"`
	Name  string              `json:"name"`
	Files []*domain.FileStore `json:"files,omitempty"`
}

type ListTypeRiceSelect struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type UpdateTypeRiceRequest struct {
	ID   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type CreateUploadFileRequest struct {
	ObjectID int64    `json:"objectId,omitempty"`
	Paths    []string `json:"paths,omitempty"`
}
