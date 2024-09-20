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
		Id:    "1",
		Email: "email",
	}, nil
}
