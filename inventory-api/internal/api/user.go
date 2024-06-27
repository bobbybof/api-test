package api

import (
	"net/http"

	"github.com/bobbybof/inventory-api/internal/helper"
	"github.com/bobbybof/inventory-api/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type createUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ValidationErrorResponse(err))
		return
	}

	hashPassword, err := helper.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ErrorHttpResponse(err, ""))
		return
	}

	user, err := server.store.CreateUser(ctx, repository.CreateUserParams{
		Name:     req.Name,
		Password: hashPassword,
		Email:    pgtype.Text{String: req.Email, Valid: len(req.Email) > 0},
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ErrorHttpResponse(err, ""))
		return
	}

	ctx.JSON(http.StatusOK, user)

}

type getUserByEmailRequest struct {
	Email string `json:"email" binding:"required"`
}

func (server *Server) GetUserByEmail(ctx *gin.Context) {
	var req getUserByEmailRequest

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ValidationErrorResponse(err))
		return
	}

	user, err := server.store.GetUserByEmail(ctx, pgtype.Text{String: req.Email, Valid: len(req.Email) > 0})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ErrorHttpResponse(err, ""))
		return
	}

	ctx.JSON(http.StatusInternalServerError, user)
}
