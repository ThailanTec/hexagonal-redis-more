package ports

import "github.com/ThailanTec/poc-serasa/internal/core/domain"

type UserRepository interface {
	CreateUser(user domain.User) (u domain.User, e error)
	GetUser(id int) (u domain.User, e error)
	DeleteUser(id int) (delete bool, e error)
}
