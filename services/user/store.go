package user

import (
	"learn/go/types"
	"log"

	"github.com/jmoiron/sqlx"
)

type UserStore interface {
	GetAll() *[]types.User
	Register(user *types.RegisterUser) (*types.User, error)
}
type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAll() *[]types.User {
	var user []types.User
	err := s.db.Select(&user, "SELECT * FROM users")
	if err != nil {
		return nil
	}
	return &user
}

func (s *Store) Register(user *types.RegisterUser) (*types.User, error) {
	tx := s.db.MustBegin()
	query := "INSERT INTO users (first_name, last_name, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING *;"

	var u types.User

	err := s.db.QueryRowx(query, user.FirstName, user.LastName, user.Email, user.PasswordHash).StructScan(&u)

	if err != nil {
		log.Println("Error in Rollback", err)
		err := tx.Rollback()
		if err != nil {
			log.Println("Another error in rollback")
			return nil, err
		}
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &u, nil
}
