package model

import (
	"time"
)

type Session struct {
	ID           string    `json:"id"`
	UserID       string    `json:"userID"`
	IpAddress    string    `json:"ipAddress"`
	UserAgent    string    `json:"userAgent"`
	LastActivity time.Time `json:"lastActivity"`
}
