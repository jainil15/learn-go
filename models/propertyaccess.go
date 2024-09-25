package models

import (
	"time"
)

type PropertyAccessStore interface {
	GetAllByUserId(userId string) ([]Property, error)
	Create(propertyId string, userId string) (*PropertyAccess, error)
}

type PropertyAccess struct {
	Id         int       `json:"id"          db:"id"`
	PropertyId int       `json:"property_id" db:"property_id"`
	UserId     int       `json:"user_id"     db:"user_id"`
	CreatedAt  time.Time `json:"created_at"  db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"  db:"updated_at"`
}
