package main

import (
	"context"
	"fmt"
	log "golang-blog-api/log"
	demo "golang-blog-api/proto"

	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	client := demo.NewFavoriteServiceClient(cc)
	request := &demo.FavoriteStatRequest{ResIds: []int32{1, 2, 3}}

	for i := 1; i <= 5; i++ {
		resp, _ := client.GetFavoriteStat(context.Background(), request)
		fmt.Printf("Receive response => [%v]\n", resp.Result)
	}
}

func init() {
	log.InitLogger(false)
}
