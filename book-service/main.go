// main.go
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hariszaki17/library-management/book-service/config"
	"github.com/hariszaki17/library-management/book-service/models"
	"github.com/hariszaki17/library-management/book-service/repository"
	"github.com/hariszaki17/library-management/book-service/usecase"
	"github.com/hariszaki17/library-management/proto/cache"
	"github.com/hariszaki17/library-management/proto/grpcclient"

	"github.com/hariszaki17/library-management/book-service/handler"

	pb "github.com/hariszaki17/library-management/proto/gen/book/proto"
	pbUser "github.com/hariszaki17/library-management/proto/gen/user/proto"

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
	db.AutoMigrate(&models.Book{})

	userConnRPC, err := grpcclient.NewGrpcConn(config.Data.UserRPCAddress)
	if err != nil {
		log.Fatalf("Failed to connect to book gRPC server: %v", err)
	}
	defer userConnRPC.Close()

	userRPC := pbUser.NewUserServiceClient(userConnRPC)

	// Instantiate Repository
	bookRepo := repository.NewBookRepository(db, redisCache)

	// Instantiate Usecase
	bookUsecase := usecase.NewBookUsecase(bookRepo, &userRPC)

	// Instantiate Handlers
	grpcHandler := handler.NewRPC(bookUsecase)

	// Start gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterBookServiceServer(grpcServer, grpcHandler)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Data.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
