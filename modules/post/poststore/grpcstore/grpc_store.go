package grpcstore

import (
	"context"
	"fmt"
	"golang-blog-api/log"
	demo "golang-blog-api/proto"

	"google.golang.org/grpc"
)

type grpcClient struct {
	client demo.FavoriteServiceClient
}

func NewGRPCClient(client demo.FavoriteServiceClient) *grpcClient {
	return &grpcClient{client: client}
}

func (c *grpcClient) GetFavoriteCountOfPosts(ctx context.Context, ids []int) (map[int]int, error) {
	resIds := make([]int32, len(ids))

	for i := range resIds {
		resIds[i] = int32(ids[i])
	}

	mapUserFavorited, err := c.client.GetFavoriteStat(ctx, &demo.FavoriteStatRequest{ResIds: resIds})

	if err != nil {
		return nil, err
	}

	result := make(map[int]int)

	for k, v := range mapUserFavorited.Result {
		result[int(k)] = int(v)
	}

	return result, nil
}

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
