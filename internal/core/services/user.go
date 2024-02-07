package services

import (
	"github.com/ThailanTec/poc-serasa/internal/core/domain"
	"github.com/ThailanTec/poc-serasa/internal/core/ports"
	"github.com/ThailanTec/poc-serasa/utils"
	erros "github.com/ThailanTec/poc-serasa/utils/errs"
)

type UserService struct {
	repo ports.UserRepository
}

func NewServices(repository ports.UserRepository) *UserService {
	return &UserService{
		repo: repository,
	}
}

func (u *UserService) CreateUser(name, email string) (domain.User, error) {
	if name == "" || email == "" {
		return domain.User{}, erros.CustomError("não foi possivel criar usuário, campo imcompleto")
	}

	nUser := domain.User{
		ID:    utils.GerenrateID(),
		Name:  name,
		Email: email,
	}

	us, err := u.repo.CreateUser(nUser)
	if err != nil {
		return domain.User{}, err
	}

	return us, nil
}

func (u *UserService) GetUserByID(id int) (domain.User, error) {

	if id == 0 {
		return domain.User{}, erros.CustomError("Não foi possivel localizar ID")
	}

	us, err := u.repo.GetUser(id)
	if err != nil {
		return domain.User{}, err
	}

	return us, nil

}

func (u *UserService) DeleteUserByID(id int) (delete bool, err error) {
	if id == 0 {
		return false, erros.CustomError("Não foi possivel localizar ID")
	}

	del, err := u.repo.DeleteUser(id)
	if err != nil {
		return false, err
	}

	return del, nil
}
