package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-rest-api/usecase"
	"github.com/gin-gonic/gin"
)

type CalculatorHandler struct {
	cu usecase.CalculatorUsecase
}

func NewCalculatorHandler(cu usecase.CalculatorUsecase) *CalculatorHandler {
	return &CalculatorHandler{
		cu: cu,
	}
}

func (h *CalculatorHandler) Add(ctx *gin.Context) {
	var numbers dto.Numbers
	err := ctx.ShouldBindJSON(&numbers)
	if err != nil {
		ctx.Error(err)
		return
	}
	res, err := h.cu.Add(ctx, numbers)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (h *CalculatorHandler) Multiply(ctx *gin.Context) {
	var numbers dto.Numbers
	err := ctx.ShouldBindJSON(&numbers)
	if err != nil {
		ctx.Error(err)
		return
	}
	res, err := h.cu.Multiply(ctx, numbers)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}
