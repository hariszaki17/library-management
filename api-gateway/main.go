package main

import (
	"fmt"
	"log"

	"github.com/hariszaki17/library-management/api-gateway/config"
	_ "github.com/hariszaki17/library-management/api-gateway/docs" // Import the generated docs package
	"github.com/hariszaki17/library-management/api-gateway/grpcclient"
	"github.com/hariszaki17/library-management/api-gateway/handler"
	"github.com/hariszaki17/library-management/api-gateway/handler/middleware"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"

	pb "github.com/hariszaki17/library-management/proto/gen/user/proto" // Replace with your proto package
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Start gRPC client connection
	userConnRPC, err := grpcclient.NewGrpcConn(cfg.UserRPCAddress)
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer userConnRPC.Close()

	userRPC := pb.NewUserServiceClient(userConnRPC)

	// Start Echo server
	e := echo.New()

	// Serve Swagger UI at /swagger/*
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Middleware to generate request ID
	e.Use(middleware.RequestIDMiddleware)

	// Serve swagger.json directly
	e.Static("/swagger", "./docs")

	// Group paths for better management
	userGroup := e.Group("/users")

	// Inject gRPC client into the handlers
	handler.NewUserHandler(userGroup, userRPC)

	// Start HTTP server
	fmt.Println("HTTP server is running on port 8080")
	log.Fatal(e.Start(fmt.Sprintf(":%s", cfg.APPPort)))
}
