package favoritestore

import (
	"context"
	demo "golang-blog-api/proto"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}

// type SQLStore interface {
// 	GetFavoriteCountOfPosts(ctx context.Context, ids []int) (map[int]int, error)
// }

type grpcStore struct {
	*sqlStore
	demo.UnimplementedFavoriteServiceServer
}

func NewGRPCStore(sqlStore *sqlStore) *grpcStore {
	return &grpcStore{
		sqlStore: sqlStore,
	}
}

func (s *grpcStore) GetFavoriteStat(ctx context.Context, request *demo.FavoriteStatRequest) (*demo.FavoriteStatResponse, error) {
	log.Println("...gRPC running GetFavoriteStat")
	resIds := make([]int, len(request.ResIds))

	for i := range resIds {
		resIds[i] = int(request.ResIds[i])
	}

	mapFavorite, err := s.GetFavoriteCountOfPosts(ctx, resIds)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "method GetFavoriteStat has something error %s", err.Error())
	}

	result := make(map[int32]int32)

	for k, v := range mapFavorite {
		result[int32(k)] = int32(v)
	}

	return &demo.FavoriteStatResponse{Result: result}, nil
}
