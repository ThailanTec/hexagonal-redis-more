package postgres

import (
	"github.com/ThailanTec/poc-serasa/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockPostgres struct {
	mock.Mock
}

func (m *MockPostgres) CreateUser(user domain.User) (u domain.User, e error) {

	args := m.Called(user)

	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockPostgres) GetUser(id int) (u domain.User, e error) {

	args := m.Called(id)

	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockPostgres) DeleteUser(id int) (delete bool, e error) {
	args := m.Called(id)

	return args.Bool(0), args.Error(1)

}
