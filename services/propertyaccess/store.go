package propertyaccess

import (
	"learn/go/models"

	"github.com/jmoiron/sqlx"
)

type PropertyAccessStore interface {
	Create(propertyId string, userId string) (*models.PropertyAccess, error)
}

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
	query := "INSERT INTO propertyaccesses ( property_id, user_id) VALUES ($1, $2) returning *"
	if err := s.db.QueryRowx(query, propertyId, userId).StructScan(&propertyaccess); err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			panic("ERROR ROLLING BACK")
		}
		return nil, err
	}
	tx.Commit()
	return &propertyaccess, nil
}
