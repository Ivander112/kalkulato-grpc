package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "Ivander112/kalkulator-grpc/rpc_function/calculator_rpc"

	"google.golang.org/grpc"
)

var serverAddr = flag.String("server", "localhost:50055", "server address")

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalcServiceClient(conn)

	var n1, n2 float32

	fmt.Print("Masukkan bilangan pertama: ")
	_, err = fmt.Scan(&n1)
	if err != nil {
		log.Fatalf("error reading operand1: %v", err)
	}

	fmt.Print("Masukkan bilangan kedua: ")
	_, err = fmt.Scan(&n2)
	if err != nil {
		log.Fatalf("error reading operand2: %v", err)
	}

	req := &pb.CalcRequest{
		Operand1: n1,
		Operand2: n2,
	}

	// Panggil metode CalcAdd di server
	addResp, err := client.CalcAdd(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling CalcAdd: %v", err)
	}
	fmt.Printf("Hasil penjumlahan: %v\n", addResp.GetResult())

	// Panggil metode CalcSubtract di server
	subtractResp, err := client.CalcSubtract(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling CalcSubtract: %v", err)
	}
	fmt.Printf("Hasil pengurangan: %v\n", subtractResp.GetResult())

	// Panggil metode CalcDivide di server
	divideResp, err := client.CalcDivide(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling CalcDivide: %v", err)
	}
	if divideResp.GetZeroDiv() {
		fmt.Println("Hasil pembagian: Pembagi 0. Pembagian tidak bisa dilakukan")
	} else {
		fmt.Printf("Hasil pembagian: %v\n", divideResp.GetResult())
	}

	// Panggil metode CalcMultiply di server
	multiplyResp, err := client.CalcMultiply(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling CalcMultiply: %v", err)
	}
	fmt.Printf("Hasil perkalian: %v\n", multiplyResp.GetResult())
}
