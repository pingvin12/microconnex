package main

import (
	"context"
	"log"
	pb "microconnex/proto"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type dateServer struct {
	pb.DateServiceServer
}

func (s *dateServer) GetEndDate(ctx context.Context, request *pb.DateRequest) (*pb.DateResponse, error) {
	// Retrieve the start date from the request
	startDate := request.GetStartDateInput()
	turnaroundTime := request.GetTurnaroundTimeNumber()

	response := date.getExpirationDate(startDate, turnaroundTime)

	response := &pb.DateResponse{
		EndDateResponse: response.Format("2006-01-02T15:04:05.000Z"),
	}

	return response, nil
}

func getExpirationDate() {
	panic("unimplemented")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server" + err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDateServiceServer(grpcServer, &dateServer{})
	log.Printf("Server started at %v", lis.Addr())

	log.Println("Starting gRPC server on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
