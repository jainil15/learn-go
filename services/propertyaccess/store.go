package propertyaccess

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

func (s *Store) Create(propertyId string, userId string) (*models.PropertyAccess, error) {
	propertyaccess := models.PropertyAccess{}
	tx := s.db.MustBegin()
	query := "INSERT INTO propertyaccesses (property_id, user_id) VALUES ($1, $2) returning *"
	if err := s.db.QueryRowx(query, propertyId, userId).StructScan(&propertyaccess); err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			panic("ERROR ROLLING BACK")
		}
		return nil, errors.New("Error creating property access")
	}
	tx.Commit()
	return &propertyaccess, nil
}

func (s *Store) GetAllByUserId(userId string) ([]models.Property, error) {
	properties := []models.Property{}
	query := `SELECT properties.* FROM properties JOIN propertyaccesses ON properties.id = propertyaccesses.property_id WHERE propertyaccesses.user_id = $1`
	if err := s.db.Select(&properties, query, userId); err != nil {
		return nil, err
	}
	return properties, nil
}
