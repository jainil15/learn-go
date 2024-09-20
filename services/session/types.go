package session

import "time"

type Session struct {
	Id        string    `json:"id"         db:"id"`
	UserId    string    `json:"user_id"    db:"user_id"`
	Expiry    time.Time `json:"-"          db:"expiry"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
