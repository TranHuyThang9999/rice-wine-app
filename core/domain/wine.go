package domain

type Wine struct {
	ID         int64   `json:"id,omitempty"`
	Name       string  `json:"name,omitempty"`
	Quantity   int     `json:"quantity,omitempty"`
	Price      float64 `json:"price,omitempty"`
	Origin     string  `json:"origin,omitempty"`      // Xuất xứ rượu
	AlcoholPct float64 `json:"alcohol_pct,omitempty"` // Nồng độ cồn
	Volume     float64 `json:"volume,omitempty"`      // Dung tích (ml, lít, v.v.)
	WineType   int     `json:"wine_type,omitempty"`   //  loại rượu
}
