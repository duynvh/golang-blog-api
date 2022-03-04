package main

import (
	"golang-blog-api/component"
	"golang-blog-api/component/uploadprovider"
	"golang-blog-api/middleware"
	"golang-blog-api/modules/category/categorytransport/gincategory"
	"golang-blog-api/modules/post/posttransport/ginpost"
	"golang-blog-api/modules/upload/uploadtransport/ginupload"
	"golang-blog-api/modules/user/usertransport/ginuser"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) error {
	appCtx := component.NewAppContext(db, upProvider, secretKey)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("v1")
	v1.POST("/upload", ginupload.Upload(appCtx))
	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx), ginuser.GetProfile(appCtx))

	categories := v1.Group("/categories", middleware.RequireAuth(appCtx))
	{
		categories.POST("", gincategory.Create(appCtx))
		categories.GET("/:id", gincategory.Get(appCtx))
		categories.GET("", gincategory.List(appCtx))
		categories.PATCH("/:id", gincategory.Update(appCtx))
		categories.DELETE("/:id", gincategory.Delete(appCtx))
	}

	posts := v1.Group("/posts", middleware.RequireAuth(appCtx))
	{
		posts.POST("", ginpost.Create(appCtx))
		posts.GET("/:id", ginpost.Get(appCtx))
		posts.GET("", ginpost.List(appCtx))
		posts.PATCH("/:id", ginpost.Update(appCtx))
		posts.DELETE("/:id", ginpost.Delete(appCtx))
	}

	return r.Run()
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Error loading .env file: ", err)
	}

	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting database: ", err)
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
		log.Fatalln("Error running service: ", err)
	}
}
