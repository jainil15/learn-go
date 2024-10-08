package user

import (
	"learn/go/models"
	"log"

	"github.com/jmoiron/sqlx"
)

type UserStore interface {
	GetAll() *[]models.User
	Register(user *models.RegisterUser) (*models.User, error)
	// GetByEmail returns a user by email
	GetByEmail(email string) (*models.User, error)
}

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetAll() *[]models.User {
	var user []models.User
	err := s.db.Select(&user, "SELECT * FROM users")
	if err != nil {
		return nil
	}
	return &user
}

func (s *Store) Register(user *models.RegisterUser) (*models.User, error) {
	tx := s.db.MustBegin()
	query := "INSERT INTO users (first_name, last_name, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING *;"

	var u models.User

	err := s.db.QueryRowx(query, user.FirstName, user.LastName, user.Email, user.PasswordHash).
		StructScan(&u)
	if err != nil {
		log.Println("Error in Rollback", err)
		txerr := tx.Rollback()
		if txerr != nil {
			log.Println("Another error in rollback")
			return nil, err
		}
		log.Println("Error in Register", err)
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *Store) GetByEmail(email string) (*models.User, error) {
	tx := s.db.MustBegin()
	query := "SELECT * FROM users WHERE email=$1 LIMIT 1"
	var u models.User
	err := s.db.Get(&u, query, email)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &u, nil
}
