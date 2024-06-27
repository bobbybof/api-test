package api

import (
	"net/http"

	"github.com/bobbybof/inventory-api/internal/helper"
	"github.com/bobbybof/inventory-api/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type getAllProductsParam struct {
	Limit  int64 `form:"limit" binding:"numeric,required"`
	Offset int64 `form:"offset" binding:"numeric"`
}

func (server *Server) GetAllProducts(ctx *gin.Context) {
	var req getAllProductsParam

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ValidationErrorResponse(err))
		return
	}

	products, totalProduct, err := server.store.GetAllProducts(ctx, repository.GetAllProductsParam{
		Limit:  req.Limit,
		Offset: req.Offset,
		Search: "",
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ErrorHttpResponse(err, ""))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  products,
		"total": totalProduct,
	})
}

type createProductParams struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
}

func (server *Server) CreateProduct(ctx *gin.Context) {
	var req createProductParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ValidationErrorResponse(err))
		return
	}

	product, err := server.store.CreateProduct(ctx, repository.CreateProductParams{
		Name:        req.Name,
		Price:       req.Price,
		Description: pgtype.Text{String: req.Description, Valid: len(req.Description) > 0},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helper.ErrorHttpResponse(err, ""))
		return
	}

	ctx.JSON(http.StatusOK, product)
}
