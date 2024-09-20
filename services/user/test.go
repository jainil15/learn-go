package user

type MockUserStore struct{}

func (m *MockUserStore) GetAll() *[]User {
	return &[]User{}
}

func (m *MockUserStore) Register(user *RegisterUser) (*User, error) {
	return nil, nil
}

func (m *MockUserStore) GetByEmail(email string) (*User, error) {
	return &User{
		Id:           "1",
		Email:        email,
		PasswordHash: "$2a$10$Ml6f3xGiiFPDZeyKH5xt2.e7FoemFtfXb4KchgD1GM5uk0kMEqHVS",
	}, nil
}
