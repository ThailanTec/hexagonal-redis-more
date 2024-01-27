package ports

import "github.com/ThailanTec/poc-serasa/internal/core/domain"

type UserService interface {
	CreateUser(name, email string) (domain.User, error)
	GetUserByID(id int) (domain.User, error)
	DeleteUserByID(id int) (delete bool, err error)
}
