package user

import "learn/go/models"

type MockUserStore struct{}

func (m *MockUserStore) GetAll() *[]models.User {
	return &[]models.User{}
}

func (m *MockUserStore) Register(user *models.RegisterUser) (*models.User, error) {
	return nil, nil
}

func (m *MockUserStore) GetByEmail(email string) (*models.User, error) {
	return &models.User{
		Id:           "1",
		Email:        email,
		PasswordHash: "$2a$10$Ml6f3xGiiFPDZeyKH5xt2.e7FoemFtfXb4KchgD1GM5uk0kMEqHVS",
	}, nil
}
