package handler

import (
	erros "github.com/ThailanTec/poc-serasa/utils/errs"
	"net/http"
	"strconv"

	"github.com/ThailanTec/poc-serasa/internal/core/domain"
	"github.com/ThailanTec/poc-serasa/internal/core/services"
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
		// Construa a resposta de erro com uma mensagem significativa
		errorResponse := erros.CustomError("Erro ao fazer Shouldbind")
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	usr, err := h.svc.CreateUser(user.Name, user.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Errooou")
		return
	}

	ctx.JSON(http.StatusCreated, usr)
}

func (h *HTTPHandler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	nId, _ := strconv.Atoi(id)

	user, err := h.svc.GetUserByID(nId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
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
