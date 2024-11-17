package entities

type CreateRiceRequest struct {
	TypeRiceID    int64    `json:"type_rice_id" binding:"required"`
	Name          string   `json:"name,omitempty" binding:"required"`
	Quantity      int      `json:"quantity,omitempty" binding:"required"`
	Price         float64  `json:"price,omitempty" binding:"required"`
	Origin        string   `json:"origin,omitempty"`         // Xuất xứ gạo
	HarvestSeason int      `json:"harvest_season,omitempty"` // Mùa thu hoạch gạo
	Files         []string `json:"files,omitempty"`
}

type ListRiceByUserIDResponse struct {
	ID            int64   `json:"id,omitempty"`
	TypeRiceID    int64   `json:"type_rice_id"`
	Name          string  `json:"name,omitempty"`
	Quantity      int     `json:"quantity,omitempty"`
	Price         float64 `json:"price,omitempty"`
	Origin        string  `json:"origin,omitempty"`         // Xuất xứ gạo
	HarvestSeason int     `json:"harvest_season,omitempty"` // Mùa thu hoạch gạo
	CreatedAt     int64   `json:"created_at,omitempty"`

	Files []string `json:"files,omitempty"`
}
