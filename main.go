package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	grpc_gen "grpc-example/gen"
	"io/ioutil"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement VehicleServer.
type server struct {
	grpc_gen.UnimplementedVehicleServiceServer
}

//  implements VehicleServer
func (s *server) GetVehicleList(ctx context.Context, request *grpc_gen.VehicleRequest) (*grpc_gen.GetVehicleResponse, error) {
	vehicles := ReadFile()

	return &grpc_gen.GetVehicleResponse{
		Vehicles: vehicles.Vehicles,
	}, nil
}

func (s *server) UpdateVehicle(ctx context.Context, request *grpc_gen.VehicleRequest) (*grpc_gen.GetVehicleResponse, error) {

	return nil, nil
}

type Vehicles struct {
	Vehicles []*grpc_gen.VehicleInfo `json:"vehicles"`
}

func ReadFile() *Vehicles {
	b, err := ioutil.ReadFile("./mock/test.json")
	if err != nil {
		fmt.Print(err)
	}

	var data Vehicles
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Fatalf("invalid format %v", err)
		return nil
	}

	return &data
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := grpc.NewServer()
	grpc_gen.RegisterVehicleServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
