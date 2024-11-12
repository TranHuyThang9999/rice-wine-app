package entities

type CreateTypeRiceRequest struct {
	Name  string   `json:"name"`
	Files []string `json:"files"`
}
