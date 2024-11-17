package domain

import (
	"context"
	"gorm.io/gorm"
)

type Rices struct {
	ID            int64   `json:"id,omitempty"`
	CreatorID     int64   `json:"creatorId,omitempty"`
	TypeRiceID    int64   `json:"type_rice_id"`
	Name          string  `json:"name,omitempty"`
	Quantity      int     `json:"quantity,omitempty"`
	Price         float64 `json:"price,omitempty"`
	Origin        string  `json:"origin,omitempty"`         // Xuất xứ gạo
	HarvestSeason int     `json:"harvest_season,omitempty"` // Mùa thu hoạch gạo
	CreatedAt     int64   `json:"created_at,omitempty"`
	UpdatedAt     int64   `json:"updated_at,omitempty"`
}
type RepositoryRice interface {
	Create(ctx context.Context, tx *gorm.DB, req *Rices) error
	GetListByCreatorID(ctx context.Context, creatorID int64) ([]*Rices, error)
	DeleteById(ctx context.Context, id int64) error
	UpdateById(ctx context.Context, req *Rices) error
	GetByRiceName(ctx context.Context, userID int64, name string) (int64, error)
}
