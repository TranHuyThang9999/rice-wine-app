package entities

type CreateTypeRiceRequest struct {
	Name  string   `json:"name" binding:"required"`
	Files []string `json:"files"`
}
type ListTypeRiceResponse struct {
	ID    int64    `json:"id"`
	Name  string   `json:"name"`
	Files []string `json:"files"`
}
