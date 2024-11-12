package domain

import (
	"context"
)

type Rice struct {
	ID            int64   `json:"id,omitempty"`
	TypeRice      int64   `json:"type_rice"`
	Name          string  `json:"name,omitempty"`
	Quantity      int     `json:"quantity,omitempty"`
	Price         float64 `json:"price,omitempty"`
	Origin        string  `json:"origin,omitempty"`         // Xuất xứ gạo
	HarvestSeason int     `json:"harvest_season,omitempty"` // Mùa thu hoạch gạo
}
type RepositoryRice interface {
	Create(ctx context.Context, req *Rice) error
	GetList(ctx context.Context) ([]*Rice, error)
	DeleteById(ctx context.Context, id int64) error
	UpdateById(ctx context.Context, req *Rice) error
}
