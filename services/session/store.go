package session

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type SessionStore interface {
	CreateSession(userId string) (*Session, error)
	GetSessionByID(userId string) (*Session, error)
}

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateSession(userId string) (*Session, error) {
	tx := s.db.MustBegin()
	query := "INSERT INTO sessions (user_id, expire) VALUES ($1, $2) RETURNING *"
	var session Session
	err := s.db.QueryRowx(query, userId, time.Hour).StructScan(&session)
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

func (s *Store) GetSessionByID(userId string) (*Session, error) {
	query := "SELECT * FROM sessions WHERE user_id=$1"
	var session Session
	err := s.db.Get(&session, query, userId)
	if err != nil {
		return nil, err
	}
	return &session, nil
}
