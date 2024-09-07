// main.go
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hariszaki17/library-management/proto/cache"
	"github.com/hariszaki17/library-management/proto/grpcclient"
	"github.com/hariszaki17/library-management/user-service/config"
	"github.com/hariszaki17/library-management/user-service/models"
	"github.com/hariszaki17/library-management/user-service/repository"
	"github.com/hariszaki17/library-management/user-service/usecase"

	"github.com/hariszaki17/library-management/user-service/handler"

	pbBook "github.com/hariszaki17/library-management/proto/gen/book/proto" // Replace with your proto package
	pb "github.com/hariszaki17/library-management/proto/gen/user/proto"     // Replace with your proto package

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
	redisCache := cache.NewCache(fmt.Sprintf("%s:%s", config.Data.RedistHost, config.Data.RedisPort))

	// Migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.BorrowingRecord{})

	bookConnRPC, err := grpcclient.NewGrpcConn(config.Data.BookRPCAddress)
	if err != nil {
		log.Fatalf("Failed to connect to book gRPC server: %v", err)
	}
	defer bookConnRPC.Close()

	bookRPC := pbBook.NewBookServiceClient(bookConnRPC)

	// Instantiate Repository
	userRepo := repository.NewUserRepository(db, redisCache)
	borrowingRecordRepo := repository.NewBorrowingRecordRepository(db, redisCache)

	// Instantiate Usecase
	userUsecase := usecase.NewUserUsecase(userRepo)
	borrowingRecordUsecase := usecase.NewBorrowingRecordUsecase(borrowingRecordRepo, &bookRPC)

	// Instantiate Handlers
	grpcHandler := handler.NewRPC(userUsecase, borrowingRecordUsecase)

	// Start gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, grpcHandler)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Data.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
