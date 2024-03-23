package main

import (
	"context"
	"fmt"
	"github.com/abbasfisal/grpc-tut/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
	"net"
)

type Server struct {
	calculatorpb.UnimplementedCalculatorServer
}

func (s *Server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {

	fmt.Println("sum request is : ", req)

	n1 := req.GetFirstNumber()
	n2 := req.GetSecNumber()

	res := n1 + n2

	response := &calculatorpb.SumResponse{Result: res}

	return response, nil

}

func (s *Server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	fmt.Println("Received Square Root ")
	number := req.GetNumber()

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("received nagative number : %v", number),
		)
	}

	return &calculatorpb.SquareRootResponse{NumberRoot: math.Sqrt(float64(number))}, nil

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
