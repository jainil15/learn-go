package propertyaccess

import (
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

func (s *Store) Create(
	propertyId string,
	userId string,
	tx *sqlx.Tx,
) (*models.PropertyAccess, error) {
	propertyaccess := models.PropertyAccess{}
	query := "INSERT INTO propertyaccesses (property_id, user_id) VALUES ($1, $2) returning *"
	if err := tx.QueryRowx(query, propertyId, userId).StructScan(&propertyaccess); err != nil {
		return nil, err
	}
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

func (s *Store) GetByUserIdAndPropertyId(
	userId string,
	propertyId string,
) (*models.Property, error) {
	property := models.Property{}
	query := `SELECT properties.* FROM properties JOIN propertyaccesses ON properties.id = propertyaccesses.property_id WHERE propertyaccesses.user_id = $1 AND propertyaccesses.property_id = $2`
	if err := s.db.Get(&property, query, userId, propertyId); err != nil {
		return nil, err
	}
	return &property, nil
}
