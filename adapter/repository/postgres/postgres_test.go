package postgres

import (
	"testing"

	"github.com/ThailanTec/poc-serasa/internal/core/domain"
	erros "github.com/ThailanTec/poc-serasa/pkg/errs"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Criar um mock do repositório Postgres
	mockRepo := new(MockPostgres)

	// Configurar comportamento esperado do mock
	expectedUser := domain.User{Name: "Test User", Email: "Thailan"}
	mockRepo.On("CreateUser", expectedUser).Return(expectedUser, nil)

	// Chamar a função CreateUser no mock
	user, err := mockRepo.CreateUser(expectedUser)

	// Verificar se não houve erro
	assert.NoError(t, err, "Erro inesperado ao criar usuário")

	// Verificar se o usuário retornado pelo mock corresponde ao esperado
	assert.Equal(t, expectedUser, user, "Usuário retornado não corresponde ao esperado")

	// Verificar se o método CreateUser foi chamado com os argumentos corretos
	mockRepo.AssertCalled(t, "CreateUser", expectedUser)
}

func TestCreateUserError(t *testing.T) {
	// Criar um mock do repositório Postgres
	mockRepo := new(MockPostgres)

	// Configurar comportamento esperado do mock para retornar um erro ao criar o usuário
	expectedUser := domain.User{Name: "Test User", Email: "thailandev@"}
	expectedError := erros.CustomError("Erro ao criar usuário")
	mockRepo.On("CreateUser", expectedUser).Return(domain.User{}, expectedError)

	// Chamar a função CreateUser no mock
	user, err := mockRepo.CreateUser(expectedUser)

	// Verificar se ocorreu o erro esperado
	assert.Error(t, err, "Esperava um erro ao criar usuário")
	assert.EqualError(t, err, expectedError.Error(), "Erro retornado não corresponde ao esperado")

	// Verificar se o usuário retornado pelo mock está vazio
	assert.Equal(t, domain.User{}, user, "O usuário retornado não deveria existir devido ao erro")

	// Verificar se o método CreateUser foi chamado com os argumentos corretos
	mockRepo.AssertCalled(t, "CreateUser", expectedUser)
}

// GET test
func TestGetUser(t *testing.T) {
	// Criar um mock do repositório Postgres
	mockRepo := new(MockPostgres)

	// Configurar comportamento esperado do mock
	expectedID := 1
	expectedUser := domain.User{ID: expectedID, Name: "Test User", Email: "thaialn"}
	mockRepo.On("GetUser", expectedID).Return(expectedUser, nil)

	// Chamar a função GetUser no mock
	user, err := mockRepo.GetUser(expectedID)

	// Verificar se não houve erro
	assert.NoError(t, err, "Erro inesperado ao buscar usuário")

	// Verificar se o usuário retornado pelo mock corresponde ao esperado
	assert.Equal(t, expectedUser, user, "Usuário retornado não corresponde ao esperado")

	// Verificar se o método GetUser foi chamado com os argumentos corretos
	mockRepo.AssertCalled(t, "GetUser", expectedID)
}

func TestDeleteUser(t *testing.T) {
	// Criar um mock do repositório Postgres
	mockRepo := new(MockPostgres)

	// Configurar comportamento esperado do mock
	expectedID := 1
	mockRepo.On("DeleteUser", expectedID).Return(true, nil)

	// Chamar a função DeleteUser no mock
	deleted, err := mockRepo.DeleteUser(expectedID)

	// Verificar se não houve erro
	assert.NoError(t, err, "Erro inesperado ao deletar usuário")

	// Verificar se o usuário foi deletado com sucesso
	assert.True(t, deleted, "Usuário não foi deletado com sucesso")

	// Verificar se o método DeleteUser foi chamado com os argumentos corretos
	mockRepo.AssertCalled(t, "DeleteUser", expectedID)
}
