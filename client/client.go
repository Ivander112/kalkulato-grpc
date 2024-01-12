package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "Ivander112/kalkulator-grpc/rpc_function/calculator_rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverAddr = flag.String("server", "localhost:50055", "The server address")

func main() {

	// Menghubungkan ke server gRPC
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Membuat instance dari CalcService client
	client := pb.NewCalcServiceClient(conn)

	// Mengambil input dari pengguna
	var n1, n2 int32

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

	// Mengisi data CalcRequest
	req := &pb.CalcRequest{
		Operand1: n1,
		Operand2: n2,
	}

	// Memanggil CalcStart di server
	resp, err := client.CalcStart(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling CalcStart: %v", err)
	}

	fmt.Printf("Hasil penjumlahan: %d\n", resp.AdditionResult)
	fmt.Printf("Hasil pengurangan: %d\n", resp.SubtractionResult)
	if resp.ZeroDiv == true {
		fmt.Printf("Hasil pembagian: Pembagi 0. Pembagian tidak bisa dilakukan\n")
	} else {
		fmt.Printf("Hasil pembagian: %f\n", resp.DivisionResult)
	}
	fmt.Printf("Hasil perkalian: %d\n", resp.MultiplicationResult)
}
