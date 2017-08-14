package main

import (
	"golang-proto/consumer/com"
	"golang-proto/consumer/logger"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/joho/godotenv"
)

const (
	port = ":50051"
)

// server is used to implement com.RouteMsgServer.
type server struct{}

// ConsumeMsg implements com.RouteMsgServer
func (s *server) SendMsg(ctx context.Context, in *com.ComMsg) (*com.StatusReport, error) {
	logger.LogMsgConsumed(*in)
	return &com.StatusReport{Status: com.StatusReport_SUCCESS, Message: in}, nil
}

func main() {
	err := godotenv.Load("general.env")
	if err != nil {
		log.Fatal("Error loading general.env file")
	}

	log.Printf("Starting gRPC Server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	com.RegisterRouteMsgServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("Server created")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
