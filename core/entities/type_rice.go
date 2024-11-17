package entities

import "rice-wine-shop/core/domain"

type CreateTypeRiceRequest struct {
	Name  string   `json:"name" binding:"required"`
	Files []string `json:"files"`
}
type ListTypeRiceResponse struct {
	ID    int64               `json:"id"`
	Name  string              `json:"name"`
	Files []*domain.FileStore `json:"files,omitempty"`
}
