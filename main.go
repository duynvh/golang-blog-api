package main

import (
	"context"
	"fmt"
	"golang-blog-api/component"
	"golang-blog-api/component/uploadprovider"
	log "golang-blog-api/log"
	"golang-blog-api/memcache"
	"golang-blog-api/middleware"
	"golang-blog-api/modules/category/categorytransport/gincategory"
	"golang-blog-api/modules/favorite/favoritestore"
	"golang-blog-api/modules/favorite/favoritetransport/ginfavorite"
	"golang-blog-api/modules/post/posttransport/ginpost"
	"golang-blog-api/modules/upload/uploadtransport/ginupload"
	"golang-blog-api/modules/user/userstore"
	"golang-blog-api/modules/user/usertransport/ginuser"
	"golang-blog-api/pubsub/pblocal"
	"golang-blog-api/skio"
	"golang-blog-api/subscriber"
	"net"
	"net/http"
	"os"

	demo "golang-blog-api/proto"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	jg "go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	log.InitLogger(false)

	if os.Getenv("APP_ENV") == "production" {

	} else {
		if err := godotenv.Load(".env.local"); err != nil {
			log.Fatal("Error loading .env file: ", err)
		}
	}
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) error {
	appCtx := component.NewAppContext(db, upProvider, secretKey, pblocal.NewPubSub())
	userStore := userstore.NewSQLStore(appCtx.GetMainDBConnection())
	userCachingStore := memcache.NewUserCaching(memcache.NewCaching(), userStore)

	r := gin.Default()

	rtEngine := skio.NewEngine()

	if err := rtEngine.Run(appCtx, r); err != nil {
		log.Fatal(err)
	}

	if err := subscriber.NewEngine(appCtx, rtEngine).Start(); err != nil {
		log.Fatal(err)
	}

	r.Use(middleware.Recover(appCtx))

	r.StaticFile("/demo/", "./demo.html")

	v1 := r.Group("v1")
	v1.POST("/upload", ginupload.Upload(appCtx))
	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx, userCachingStore), ginuser.GetProfile(appCtx))
	v1.GET("/favorited-posts", middleware.RequireAuth(appCtx, userCachingStore), ginfavorite.ListFavoritedPostsOfAUser(appCtx))

	categories := v1.Group("/categories", middleware.RequireAuth(appCtx, userCachingStore))
	{
		categories.POST("", gincategory.Create(appCtx))
		categories.GET("/:id", gincategory.Get(appCtx))
		categories.GET("", gincategory.List(appCtx))
		categories.PATCH("/:id", gincategory.Update(appCtx))
		categories.DELETE("/:id", gincategory.Delete(appCtx))
	}

	posts := v1.Group("/posts", middleware.RequireAuth(appCtx, userCachingStore))
	{
		posts.POST("", ginpost.Create(appCtx))
		posts.GET("/:id", ginpost.Get(appCtx))
		posts.GET("", ginpost.List(appCtx))
		posts.PATCH("/:id", ginpost.Update(appCtx))
		posts.DELETE("/:id", ginpost.Delete(appCtx))
		posts.POST("/:id/favorite", ginfavorite.Favorite(appCtx))
		posts.DELETE("/:id/unfavorite", ginfavorite.Unfavorite(appCtx))
		posts.GET("/:id/favorited-users", ginfavorite.ListUsersFavoritedAPost(appCtx))
	}

	je, err := jg.NewExporter(jg.Options{
		AgentEndpoint: "localhost:6831",
		Process:       jg.Process{ServiceName: "Golang-blog-API"},
	})

	if err != nil {
		log.Print(err)
	}

	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(1)})

	// setupGRPCServer(appCtx, db)
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	fmt.Printf("Server is running on %v ...\n", address)

	s := grpc.NewServer()
	gRPCStore := favoritestore.NewGRPCStore(favoritestore.NewSQLStore(db))
	demo.RegisterFavoriteServiceServer(s, gRPCStore)

	go s.Serve(lis)

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connect to gRPC server successful")

	defer cc.Close()
	appCtx.SetGRPCClientConnection(cc)

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:50051",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = demo.RegisterFavoriteServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatal("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Print("Serving gRPC-Gateway on http://0.0.0.0:8090")
	go gwServer.ListenAndServe()

	return http.ListenAndServe(
		":8080",
		&ochttp.Handler{
			Handler: r,
		},
	)
}

func main() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting database: ", err)
	}

	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(
		os.Getenv("S3_BUCKET_NAME"),
		os.Getenv("S3_REGION"),
		os.Getenv("S3_API_KEY"),
		os.Getenv("S3_SECRET_KEY"),
		os.Getenv("S3_DOMAIN"),
	)

	secretKey := os.Getenv("SYSTEM_SECRET")

	if err := runService(db, s3Provider, secretKey); err != nil {
		log.Fatal("Error running service: ", err)
	}
}

// func setupGRPCServer(appCtx component.AppContext, db *gorm.DB) {
// 	address := "0.0.0.0:50051"
// 	lis, err := net.Listen("tcp", address)

// 	if err != nil {
// 		log.Fatalf("Error %v", err)
// 	}

// 	fmt.Printf("Server is running on %v ...\n", address)

// 	s := grpc.NewServer()
// 	gRPCStore := favoritestore.NewGRPCStore(favoritestore.NewSQLStore(db))
// 	demo.RegisterFavoriteServiceServer(s, gRPCStore)

// 	go s.Serve(lis)

// 	opts := grpc.WithInsecure()
// 	cc, err := grpc.Dial("localhost:50051", opts)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Print("Connect to gRPC server successful")

// 	defer cc.Close()
// 	appCtx.SetGRPCClientConnection(cc)
// }
