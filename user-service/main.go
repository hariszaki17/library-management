// main.go
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/hariszaki17/library-management/user-service/models"
	"github.com/hariszaki17/library-management/user-service/repository"
	"github.com/hariszaki17/library-management/user-service/usecase"

	"github.com/hariszaki17/library-management/user-service/handler"

	pb "github.com/hariszaki17/library-management/proto/gen/user/proto" // Replace with your proto package

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	// Initialize GORM DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	// Instantiate Repository
	userRepo := repository.NewUserRepository(db)

	// Instantiate Usecase
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Instantiate Handlers
	grpcHandler := handler.NewGrpcHandler(userUsecase)
	// httpHandler := handler.NewHttpHandler(userUsecase)

	// Start gRPC server
	go func() {
		grpcServer := grpc.NewServer()
		pb.RegisterUserServiceServer(grpcServer, grpcHandler)

		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		fmt.Println("gRPC server is running on port 50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Start HTTP server
	// http.HandleFunc("/user", httpHandler.GetUser)
	fmt.Println("HTTP server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
