package property

import (
	"errors"
	"learn/go/models"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Create(p *models.CreatePropertyPayload) (*models.Property, error) {
	property := models.Property{}
	tx := s.db.MustBegin()
	query := "INSERT INTO properties ( name, email, phone_number, address, about) VALUES ($1, $2, $3, $4, $5) returning *"
	if err := tx.QueryRowx(query, p.Name, p.Email, p.PhoneNumber, p.Address, p.About).StructScan(&property); err != nil {
		tx.Rollback()
		return nil, errors.New("Error creating property")
	}
	tx.Commit()
	return &property, nil
}
