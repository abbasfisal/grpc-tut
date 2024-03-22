package main

import (
	"context"
	"fmt"
	"github.com/abbasfisal/grpc-tut/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	//client
	c := calculatorpb.NewCalculatorClient(cc)

	req := calculatorpb.SumRequest{
		FirstNumber: 20,
		SecNumber:   80,
	}
	response, err := c.Sum(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("sum result is : ", response)

}
