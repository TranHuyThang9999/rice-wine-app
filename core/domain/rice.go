package domain

import (
	"context"
)

type Rice struct {
	ID            int64   `json:"id,omitempty"`
	Name          string  `json:"name,omitempty"`
	Quantity      int     `json:"quantity,omitempty"`
	Price         float64 `json:"price,omitempty"`
	Origin        string  `json:"origin,omitempty"`         // Xuất xứ gạo
	RiceType      int     `json:"rice_type,omitempty"`      // Loại gạo (nếp, tẻ, v.v.)
	HarvestSeason int     `json:"harvest_season,omitempty"` // Mùa thu hoạch gạo

}
type RepositoryRice interface {
	Create(ctx context.Context, req *Rice) error
	GetList(ctx context.Context) ([]*Rice, error)
	DeleteById(ctx context.Context, id int64) error
	UpdateById(ctx context.Context, req *Rice) error
}
