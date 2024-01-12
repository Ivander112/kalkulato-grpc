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

var port = flag.Int("port", 50055, "port server")

type server struct {
	pb.UnimplementedCalcServiceServer
}

func (s *server) CalcAdd(ctx context.Context, req *pb.CalcRequest) (*pb.CalcAddResponse, error) {
	result := req.GetOperand1() + req.GetOperand2()
	return &pb.CalcAddResponse{Result: result}, nil
}

func (s *server) CalcSubtract(ctx context.Context, req *pb.CalcRequest) (*pb.CalcSubtractResponse, error) {
	result := req.GetOperand1() - req.GetOperand2()
	return &pb.CalcSubtractResponse{Result: result}, nil
}

func (s *server) CalcDivide(ctx context.Context, req *pb.CalcRequest) (*pb.CalcDivideResponse, error) {
	operand2 := req.GetOperand2()
	var zeroDiv bool
	var result float32

	if operand2 == 0 {
		zeroDiv = true
		result = 0
	} else {
		zeroDiv = false
		result = req.GetOperand1() / operand2
	}

	return &pb.CalcDivideResponse{Result: result, ZeroDiv: zeroDiv}, nil
}

func (s *server) CalcMultiply(ctx context.Context, req *pb.CalcRequest) (*pb.CalcMultiplyResponse, error) {
	result := req.GetOperand1() * req.GetOperand2()
	return &pb.CalcMultiplyResponse{Result: result}, nil
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
