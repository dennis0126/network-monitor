package model

import (
	"time"
)

type User struct {
	ID           string `json:"id"`
	Name         string
	PasswordHash string
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time
}
