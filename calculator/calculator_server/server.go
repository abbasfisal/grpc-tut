package main

import (
	"context"
	"fmt"
	"github.com/abbasfisal/grpc-tut/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	calculatorpb.UnimplementedCalculatorServer
}

func (s Server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {

	fmt.Println("sum request is : ", req)

	n1 := req.GetFirstNumber()
	n2 := req.GetSecNumber()

	res := n1 + n2

	response := &calculatorpb.SumResponse{Result: res}

	return response, nil

}
func main() {

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServer(s, &Server{})

	fmt.Println("calculator server is running on port :50051")
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
