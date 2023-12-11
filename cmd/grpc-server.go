package cmd

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-rest-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-rest-api/pb"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/exercise-go-grpc-rest-api/usecase"
	"google.golang.org/grpc"
)

func StartGRPCServer() {
	cu := usecase.NewCaculatorUsecase()
	ch := handler.NewCalculatorGrpcHandler(cu)

	list, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Print("error starting tcp server")
	}

	server := grpc.NewServer()
	pb.RegisterArithmeticServer(server, ch)

	log.Print("starting grpc server")

	signCh := make(chan os.Signal, 1)
	signal.Notify(signCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.Serve(list); err != nil {
			signCh <- syscall.SIGINT
		}
	}()
	log.Print("server started")
	<-signCh
	log.Print("server stopped")
}
