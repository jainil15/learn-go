package session

import "time"

type MockStore struct{}

func (m *MockStore) CreateSession(userId string) (*Session, error) {
	return &Session{
		UserId:    "1",
		Id:        "1",
		ExpiresAt: time.Now(),
	}, nil
}

func (m *MockStore) GetSessionByID(userId string) (*Session, error) {
	return &Session{
		UserId:    "1",
		Id:        "1",
		ExpiresAt: time.Now(),
	}, nil
}
