package main

import (
	"context"
	"fmt"
	"github.com/abbasfisal/grpc-tut/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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

	//doUnary(c)

	doErrUnary(c)
}

func doErrUnary(c calculatorpb.CalculatorClient) {
	fmt.Println("start unary square root\n")

	// correct call
	doErrCall(c, 30)

	// negative call
	doErrCall(c, -10)

}

func doErrCall(c calculatorpb.CalculatorClient, number int) {
	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: int32(number)})
	if err != nil {
		resErr, ok := status.FromError(err)
		if ok {
			// user error
			fmt.Println(resErr.Message())
			fmt.Println(resErr.Code())
			if resErr.Code() == codes.InvalidArgument {
				fmt.Println("negative number is not acceptable\n")
				return
			}
		} else {
			log.Fatal(err)
			return
		}
	}
	fmt.Printf("result of square : %v : %v \n", number, res.GetNumberRoot())
}

func doUnary(c calculatorpb.CalculatorClient) {

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
