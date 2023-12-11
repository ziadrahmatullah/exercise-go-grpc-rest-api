package handler

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-rest-api/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-rest-api/usecase"
)

type CalculatorGrpcHandler struct {
	pb.UnimplementedArithmeticServer
	usecase usecase.CalculatorUsecase
}

func NewAuthGrpcHandler(usecase usecase.CalculatorUsecase) *CalculatorGrpcHandler {
	return &CalculatorGrpcHandler{
		usecase: usecase,
	}
}

func (h *CalculatorGrpcHandler) Add(ctx context.Context, req *pb.Numbers)(*pb.Result, error){
	res, err := h.usecase.Add(ctx, dto.Numbers{
		Num1: int(req.Num1),
		Num2: int(req.Num2),
	})
	if err != nil {
		return nil, err
	}
	return &pb.Result{Value: int32(res.Value)}, nil
}

func (h *CalculatorGrpcHandler) Multiply(ctx context.Context, req *pb.Numbers)(*pb.Result, error){
	res, err := h.usecase.Add(ctx, dto.Numbers{
		Num1: int(req.Num1),
		Num2: int(req.Num2),
	})
	if err != nil {
		return nil, err
	}
	return &pb.Result{Value: int32(res.Value)}, nil
}
