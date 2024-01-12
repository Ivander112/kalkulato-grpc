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

// Fungsi penambahan
func (s *server) CalcAdd(ctx context.Context, req *pb.CalcRequest) (*pb.CalcAddResponse, error) {
	log.Printf("User melakukan request kalkulasi dengan Operand1: %v, Operand2: %v", req.GetOperand1(), req.GetOperand2())
	result := req.GetOperand1() + req.GetOperand2()
	return &pb.CalcAddResponse{Result: result}, nil
}

// Fungsi pengurangan
func (s *server) CalcSubtract(ctx context.Context, req *pb.CalcRequest) (*pb.CalcSubtractResponse, error) {
	result := req.GetOperand1() - req.GetOperand2()
	return &pb.CalcSubtractResponse{Result: result}, nil
}

// Fungsi pembagian
func (s *server) CalcDivide(ctx context.Context, req *pb.CalcRequest) (*pb.CalcDivideResponse, error) {
	operand2 := req.GetOperand2()
	var zeroDiv bool
	var result float32
	// Check apakah pembagi nya 0
	if operand2 == 0 {
		zeroDiv = true
		result = 0
	} else {
		zeroDiv = false
		result = req.GetOperand1() / operand2
	}

	return &pb.CalcDivideResponse{Result: result, ZeroDiv: zeroDiv}, nil
}

// Fungsi perkalian
func (s *server) CalcMultiply(ctx context.Context, req *pb.CalcRequest) (*pb.CalcMultiplyResponse, error) {
	result := req.GetOperand1() * req.GetOperand2()
	return &pb.CalcMultiplyResponse{Result: result}, nil
}

func main() {
	flag.Parse()
	// Mengatur koneksi dan port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Membuat server gRPC
	s := grpc.NewServer()
	pb.RegisterCalcServiceServer(s, &server{})
	fmt.Printf("Server is listening on port %d...\n", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
