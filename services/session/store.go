package session

import (
	"learn/go/models"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type SessionStore interface {
	CreateSession(userId string) (*models.Session, error)
	GetSessionByID(userId string) (*models.Session, error)
}

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateSession(userId string) (*models.Session, error) {
	log.Printf("Creating session for user: %v", userId)
	tx := s.db.MustBegin()
	query := "INSERT INTO sessions (user_id, expires_at) VALUES ($1, $2) RETURNING *"
	var session models.Session
	err := s.db.QueryRowx(query, userId, time.Now()).StructScan(&session)
	if err != nil {
		errR := tx.Rollback()
		if errR != nil {
			return nil, errR
		}
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		errR := tx.Rollback()
		if errR != nil {
			return nil, errR
		}
		return nil, err
	}
	return &session, nil
}

func (s *Store) GetSessionByID(userId string) (*models.Session, error) {
	log.Printf("Getting session for user: %v", userId)
	query := "SELECT * FROM sessions WHERE user_id=$1"
	var session models.Session
	err := s.db.Get(&session, query, userId)
	if err != nil {
		return nil, err
	}
	return &session, nil
}
