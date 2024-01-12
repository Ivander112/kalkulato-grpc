package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "Ivander112/kalkulator-grpc/rpc_function/calculator_rpc"
)
// alamat server
var serverAddr = flag.String("server", "localhost:50055", "alamat server")

func main() {
	flag.Parse()
	// Melakukan koneksi ke server
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalcServiceClient(conn)
	// Input angka dari user
	var n1, n2 float32
	// memastikan operand1 adalah angka
	input1:
	fmt.Print("Masukkan bilangan pertama: ")
	_, err = fmt.Scan(&n1)
	if err != nil {
		log.Printf("Bilangan harus angka: %v", err)
		goto input1
	}
	// memastikan operand2 adalah angka
	input2:
	fmt.Print("Masukkan bilangan kedua: ")
	_, err = fmt.Scan(&n2)
	if err != nil {
		log.Printf("Bilangan harus angka: %v", err)
		goto input2
	}

	req := &pb.CalcRequest{
		Operand1: n1,
		Operand2: n2,
	}

	// Memanggil fungsi CalcAdd dari server
	addResp, err := client.CalcAdd(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling CalcAdd: %v", err)
	}
	fmt.Printf("Hasil penjumlahan: %v\n", addResp.GetResult())

	// Memanggil fungsi CalcSubtract dari server
	subtractResp, err := client.CalcSubtract(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling CalcSubtract: %v", err)
	}
	fmt.Printf("Hasil pengurangan: %v\n", subtractResp.GetResult())

	// Memanggil fungsi CalcDivide dari server
	divideResp, err := client.CalcDivide(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling CalcDivide: %v", err)
	}
	if divideResp.GetZeroDiv() {
		fmt.Println("Hasil pembagian: Pembagi 0. Pembagian tidak bisa dilakukan")
	} else {
		fmt.Printf("Hasil pembagian: %v\n", divideResp.GetResult())
	}

	// Memanggil metode CalcMultiply dari server
	multiplyResp, err := client.CalcMultiply(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling CalcMultiply: %v", err)
	}
	fmt.Printf("Hasil perkalian: %v\n", multiplyResp.GetResult())
}
