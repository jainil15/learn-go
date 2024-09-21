package session

import (
	"learn/go/models"
	"time"
)

type MockStore struct{}

func (m *MockStore) CreateSession(userId string) (*models.Session, error) {
	return &models.Session{
		UserId:    "1",
		Id:        "1",
		ExpiresAt: time.Now(),
	}, nil
}

func (m *MockStore) GetSessionByID(userId string) (*models.Session, error) {
	return &models.Session{
		UserId:    "1",
		Id:        "1",
		ExpiresAt: time.Now(),
	}, nil
}
