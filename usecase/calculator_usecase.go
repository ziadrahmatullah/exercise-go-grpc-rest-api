package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-rest-api/dto"
)

type CalculatorUsecase interface {
	Add(context.Context, dto.Numbers)(*dto.Result, error)
	Multiply(context.Context, dto.Numbers)(*dto.Result, error)
}

type calculatorUsecase struct{}

func NewCaculatorUsecase() *calculatorUsecase{
	return &calculatorUsecase{}
}

func (cu *calculatorUsecase) Add(ctx context.Context, req dto.Numbers) (*dto.Result, error) {
	return &dto.Result{
		Value: (req.Num1 + req.Num2),
	}, nil
}

func (cu *calculatorUsecase) Multiply(ctx context.Context, req dto.Numbers) (*dto.Result, error) {
	return &dto.Result{
		Value: (req.Num1 * req.Num2),
	}, nil
}
