// main.go

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "Ivander112/kalkulator-grpc/rpc_function/calculator_rpc"

	"google.golang.org/grpc"
)

var port = flag.Int("port", 50055, "The server port")

type server struct {
	pb.UnimplementedCalcServiceServer
}

func (s *server) CalcStart(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	n1 := req.GetOperand1()
	n2 := req.GetOperand2()

	add := n1 + n2
	sub := n1 - n2
	mult := n2 * n2

	var zeroDiv bool
	var divisionResult float32

	if n2 == 0 {
		zeroDiv = true
		divisionResult = 0
	} else {
		zeroDiv = false
		divisionResult = float32(n1 / n2)
	}

	response := &pb.CalcResponse{
		AdditionResult:       int32(add),
		SubtractionResult:    sub,
		DivisionResult:       divisionResult,
		MultiplicationResult: mult,
		ZeroDiv:              zeroDiv,
	}

	return response, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalcServiceServer(s, &server{})
	fmt.Printf("Server is listening on port %d...\n", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
