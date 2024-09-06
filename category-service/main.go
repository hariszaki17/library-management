// main.go
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hariszaki17/library-management/proto/cache"
	"github.com/hariszaki17/library-management/category-service/config"
	"github.com/hariszaki17/library-management/category-service/models"
	"github.com/hariszaki17/library-management/category-service/repository"
	"github.com/hariszaki17/library-management/category-service/usecase"

	"github.com/hariszaki17/library-management/category-service/handler"

	pb "github.com/hariszaki17/library-management/proto/gen/category/proto"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	err := config.Load(".env")
	if err != nil {
		panic(err)
	}

	// Initialize GORM DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Data.DBHost, config.Data.DBUser, config.Data.DBPassword, config.Data.DBName, config.Data.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Initialize Redis cache
	redisCache := cache.NewCache(fmt.Sprintf("%s:%s", config.Data.RedistHost, config.Data.RedistPort))

	// Migrate the schema
	db.AutoMigrate(&models.Category{})

	// Instantiate Repository
	categoryRepo := repository.NewCategoryRepository(db, redisCache)

	// Instantiate Usecase
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)

	// Instantiate Handlers
	grpcHandler := handler.NewRPC(categoryUsecase)

	// Start gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, grpcHandler)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Data.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
