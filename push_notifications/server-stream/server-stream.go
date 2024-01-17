package main

import (
	// "fmt"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	pb "Ivander112/kalkulator-grpc/push_notifications/rpc-stream"

	"google.golang.org/grpc"
)

// port server
var port = flag.Int("port", 50055, "port server")

type NotificationsServer struct {
	pb.UnimplementedNotificationsServiceServer
}

// Struktur data untuk objek JSON
type Fruit struct {
	Family    string `json:"family"`
	Genus     string `json:"genus"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Nutrients struct {
		Calories      float32 `json:"calories"`
		Carbohydrates float32 `json:"carbohydrates"`
		Fat           float32 `json:"fat"`
		Protein       float32 `json:"protein"`
		Sugar         float32 `json:"sugar"`
	} `json:"nutritions"`
	Order string `json:"order"`
}

// Struktur data untuk objek JSON universitas
type University struct {
	Name          string   `json:"name"`
	AlphaTwoCode  string   `json:"alpha_two_code"`
	StateProvince string   `json:"state-province"`
	Domains       []string `json:"domains"`
	Country       string   `json:"country"`
	Web_Pages      []string   `json:"web_pages"`
}

// getDataFromURI mengambil data JSON dari URI dan mengembalikan array dari struktur yang sesuai
func getDataFromURI(uri string, target interface{}) error {
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &target); err != nil {
		return err
	}

	return nil
}

func (s *NotificationsServer) FruitsNotifications(req *pb.NotificationsRequest, stream pb.NotificationsService_FruitsNotificationsServer) error {
	log.Printf("Users mengikuti layanan notifikasi nutrisi pada buah")

	// Mengambil data dari URI untuk buah
	fruitURI := "https://www.fruityvice.com/api/fruit/all"
	var fruits []Fruit

	if err := getDataFromURI(fruitURI, &fruits); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(fruits)
	for _, fruit := range fruits {
		response := &pb.FruitResponse{
			Name:          fruit.Name,
			Calories:      float32(fruit.Nutrients.Calories),
			Carbohydrates: float32(fruit.Nutrients.Carbohydrates),
			Fat:           float32(fruit.Nutrients.Fat),
			Protein:       float32(fruit.Nutrients.Protein),
			Sugar:         float32(fruit.Nutrients.Sugar),
		}

		if err := stream.Send(response); err != nil {
			log.Printf("Error sending data to stream: %v", err)
			return err
		}

		// Simulasi penundaan untuk setiap buah yang dikirim
		time.Sleep(200 * time.Millisecond)
	}

	return nil
}

func (s *NotificationsServer) UniversitiesNotifications(req *pb.NotificationsRequest, stream pb.NotificationsService_UniversitiesNotificationsServer) error {
	log.Printf("Users mengikuti layanan notifikasi daftar universitas")
	country := req.Notification_Name
	// Memastikan country tidak boleh kosong
	if len(country) == 0 {
		log.Fatalln("Nama negara tidak boleh kosong")
	}

	// Mengambil data dari URI untuk universitas
	universityURI := "http://universities.hipolabs.com/search?country="+country
	log.Printf(universityURI)
	
	var universities []University
	if err := getDataFromURI(universityURI, &universities); err != nil {
		log.Fatal(err)
		return err
	}
	// Jika negara yang dicari tidak ada
	if len(universities) == 0 {
		log.Fatal("Tidak ada universitas untuk negara tersebut")
	}	
	for _, university := range universities {
		webPages := ""
		if len(university.Web_Pages) > 0 {
			webPages = university.Web_Pages[0]
		}
		response := &pb.UniversitiesResponse{
			Name:      university.Name,
			Web_Pages: webPages,
		}
		// fmt.Println(response)
		if err := stream.Send(response); err != nil {
			log.Printf("Error sending data to stream: %v", err)
			return err
		}

		// Simulasi penundaan untuk setiap universitas yang dikirim
		time.Sleep(300 * time.Millisecond)
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterNotificationsServiceServer(grpcServer, &NotificationsServer{})

	log.Printf("Starting Push Notification server on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 50051: %v", err)
	}
}
