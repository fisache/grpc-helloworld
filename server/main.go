package main

import (
	"golang.org/x/net/context"
	pb "github.com/fisache/grpc-helloworld/helloworld"
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
  "fmt"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message:"Hello " + in.Name}, nil
}

func main() {
  fmt.Println("staring server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
