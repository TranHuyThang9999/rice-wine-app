package entities

type CreateUsersRequest struct {
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Avatar      string `json:"avatar"`
}
