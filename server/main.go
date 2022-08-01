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

type myType []map[string]string

// server is used to implement VehicleServer.
type server struct {
	grpc_gen.UnimplementedVehicleServiceServer
}

//  implements VehicleServer
func (s *server) GetVehicleList(ctx context.Context, request *grpc_gen.VehicleRequest) (*grpc_gen.GetVehicleResponse, error) {
	vehicles := ReadFile()

	return &grpc_gen.GetVehicleResponse{
		Vehicles: vehicles,
	}, nil
}

func (s *server) UpdateVehicle(ctx context.Context, request *grpc_gen.VehicleRequest) (*grpc_gen.GetVehicleResponse, error) {

	return nil, nil
}

func ReadFile() []*grpc_gen.VehicleInfo {
	b, err := ioutil.ReadFile("./test.json")
	if err != nil {
		fmt.Print(err)
	}

	var data []*grpc_gen.VehicleInfo
	//type data []map[string]string

	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil
	}

	return nil
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
