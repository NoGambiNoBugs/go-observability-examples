package entity

import "github.com/google/uuid"

// Customer contains the data of customer.
type Customer struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
