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

func (s *Store) Create(p *models.CreatePropertyPayload, tx *sqlx.Tx) (*models.Property, error) {
	property := models.Property{}
	query := "INSERT INTO properties ( name, email, phone_number, address, about) VALUES ($1, $2, $3, $4, $5) returning *"
	if err := tx.QueryRowx(query, p.Name, p.Email, p.PhoneNumber, p.Address, p.About).StructScan(&property); err != nil {
		return nil, errors.New("Error creating property")
	}
	return &property, nil
}

func (s *Store) GetById(Id string) (*models.Property, error) {
	query := "SELECT * FROM properties WHERE id=$1"
	var property models.Property
	err := s.db.Get(&property, query, Id)
	if err != nil {
		return nil, err
	}
	return &property, nil
}
