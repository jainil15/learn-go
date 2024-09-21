package models

import "time"

type Session struct {
	Id        string    `json:"id"         db:"id"`
	UserId    string    `json:"user_id"    db:"user_id"`
	ExpiresAt time.Time `json:"-"          db:"expires_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
