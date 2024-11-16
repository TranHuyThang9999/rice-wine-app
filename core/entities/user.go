package entities

type CreateUsersRequest struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password" binding:"required"`
	Avatar      string `json:"avatar"`
}
type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email,omitempty"`
	Avatar      string `json:"avatar"`
	Role        int    `json:"role"`
	UpdatedAt   int64  `json:"updated_at,omitempty"`
}
