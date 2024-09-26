package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type PropertyAccessStore interface {
	GetAllByUserId(userId string) ([]Property, error)
	GetByUserIdAndPropertyId(userId string, propertyId string) (*Property, error)
	Create(propertyId string, userId string, tx *sqlx.Tx) (*PropertyAccess, error)
}

type PropertyAccess struct {
	Id         string    `json:"id"          db:"id"`
	PropertyId string    `json:"property_id" db:"property_id"`
	UserId     string    `json:"user_id"     db:"user_id"`
	CreatedAt  time.Time `json:"created_at"  db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"  db:"updated_at"`
}
