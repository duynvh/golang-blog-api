package main

import (
	"context"
	"fmt"
	log "golang-blog-api/log"
	"net"

	"google.golang.org/grpc"

	demo "golang-blog-api/proto"
)

type server struct {
	demo.UnimplementedFavoriteServiceServer
}

func (s *server) GetFavoriteStat(ctx context.Context, req *demo.FavoriteStatRequest) (*demo.FavoriteStatResponse, error) {
	log.Print("Server is received:", req.ResIds)

	return &demo.FavoriteStatResponse{
		Result: map[int32]int32{1: 1, 2: 4},
	}, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	fmt.Printf("Server is running on %v ...", address)

	s := grpc.NewServer()

	demo.RegisterFavoriteServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func init() {
	log.InitLogger(false)
}