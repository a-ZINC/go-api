package model

type Student struct {
	ID int `jsone:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Phone string `json:"phone,omitempty"`
	Address string `json:"address,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Class string `json:"class"`
}