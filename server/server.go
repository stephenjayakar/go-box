package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/stephenjayakar/go-box/gobox"
	grpc "google.golang.org/grpc"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type goBoxServer struct {
	pb.UnimplementedGoBoxServer
}

func (s *goBoxServer) Meow(ctx context.Context, request *pb.MeowRequest) (*pb.MeowResponse, error) {
	name := request.Name
	return &pb.MeowResponse{Meow: fmt.Sprintf("meow %s", name)}, nil
}

func newServer() *goBoxServer {
	s := &goBoxServer{}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGoBoxServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
