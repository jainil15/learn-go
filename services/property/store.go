package property

import (
	"learn/go/models"

	"github.com/jmoiron/sqlx"
)

type PropertyStore interface {
	Create(p *models.CreatePropertyPayload) (*models.Property, error)
}

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
	if err := s.db.QueryRowx(query, p.Name, p.Email, p.PhoneNumber, p.Address, p.About).StructScan(&property); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &property, nil
}
