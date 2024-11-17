package domain

type OrderItems struct {
	ID        int64   `json:"id"`
	OrderID   int64   `json:"order_id"`
	ProductID int64   `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Total     float64 `json:"total"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
}

type Orders struct {
	ID          int64   `json:"id"`
	UserID      int64   `json:"user_id"`
	Status      string  `json:"status"`
	TotalAmount float64 `json:"total_amount"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
}
