package handler

import (
	"net/http"
	"strconv"

	"github.com/ThailanTec/poc-serasa/internal/core/domain"
	"github.com/ThailanTec/poc-serasa/internal/core/services"
	erros "github.com/ThailanTec/poc-serasa/pkg/errs"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	svc services.UserService
}

func NewHttpHandler(svc services.UserService) *HTTPHandler {
	return &HTTPHandler{
		svc: svc,
	}
}

func (h *HTTPHandler) CreateUser(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})

		return
	}

	usr, err := h.svc.CreateUser(user.Name, user.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, erros.BadRequest("Não foi possivel criar usuário"))
		return
	}

	ctx.JSON(http.StatusCreated, usr)
}

func (h *HTTPHandler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	nId, _ := strconv.Atoi(id)

	user, err := h.svc.GetUserByID(nId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, erros.BadRequest("Não foi possivel localizar usuários"))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *HTTPHandler) DeleteUserByID(ctx *gin.Context) {

	id := ctx.Param("id")
	nId, _ := strconv.Atoi(id)

	deleted, err := h.svc.DeleteUserByID(nId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, erros.BadRequest("Erro ao deletar usuário"))
		return
	}

	if !deleted {
		ctx.JSON(http.StatusBadRequest, erros.BadRequest("Não foi possivel deletar usuário"))
		return
	} else {
		ctx.JSON(http.StatusOK, "Deletado com sucesso nengue!")
	}
}
