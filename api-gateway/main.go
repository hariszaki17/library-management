package main

import (
	"fmt"
	"log"

	"github.com/hariszaki17/library-management/api-gateway/config"
	_ "github.com/hariszaki17/library-management/api-gateway/docs" // Import the generated docs package
	"github.com/hariszaki17/library-management/api-gateway/handler"
	"github.com/hariszaki17/library-management/api-gateway/handler/middleware"
	"github.com/hariszaki17/library-management/proto/grpcclient"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	pbAuthor "github.com/hariszaki17/library-management/proto/gen/author/proto"     // Replace with your proto package
	pbBook "github.com/hariszaki17/library-management/proto/gen/book/proto"         // Replace with your proto package
	pbCategory "github.com/hariszaki17/library-management/proto/gen/category/proto" // Replace with your proto package
	pbUser "github.com/hariszaki17/library-management/proto/gen/user/proto"         // Replace with your proto package
	// Replace with your proto package
)

func main() {
	// Load configuration
	err := config.Load(".env")
	if err != nil {
		panic(err)
	}

	// Start gRPC client connection
	userConnRPC, err := grpcclient.NewGrpcConn(config.Data.UserRPCAddress)
	if err != nil {
		log.Fatalf("Failed to connect to user gRPC server: %v", err)
	}
	defer userConnRPC.Close()

	bookConnRPC, err := grpcclient.NewGrpcConn(config.Data.BookRPCAddress)
	if err != nil {
		log.Fatalf("Failed to connect to book gRPC server: %v", err)
	}
	defer bookConnRPC.Close()

	authorConnRPC, err := grpcclient.NewGrpcConn(config.Data.AuthorRPCAddress)
	if err != nil {
		log.Fatalf("Failed to connect to author gRPC server: %v", err)
	}
	defer bookConnRPC.Close()

	categoryConnRPC, err := grpcclient.NewGrpcConn(config.Data.CategoryRPCAddress)
	if err != nil {
		log.Fatalf("Failed to connect to category gRPC server: %v", err)
	}
	defer categoryConnRPC.Close()

	userRPC := pbUser.NewUserServiceClient(userConnRPC)
	bookRPC := pbBook.NewBookServiceClient(bookConnRPC)
	authorRPC := pbAuthor.NewAuthorServiceClient(authorConnRPC)
	categoryRPC := pbCategory.NewCategoryServiceClient(categoryConnRPC)

	// Start Echo server
	e := echo.New()

	// Serve Swagger UI at /swagger/*
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Middleware to generate request ID
	e.Use(middleware.RequestIDMiddleware)

	// Serve swagger.json directly
	e.Static("/swagger", "./docs")

	authMiddleware := middleware.AuthJWTMiddleware(userRPC)

	// Group paths for better management
	userGroup := e.Group("/users")
	authGroup := e.Group("/auth")
	bookGroup := e.Group("/books")
	authorGroup := e.Group("/authors")
	categoryGroup := e.Group("/categories")
	borrowingRecordGroup := e.Group("/borrow-book")

	// Inject gRPC client into the handlers
	handler.NewAuthHandler(authGroup, userRPC)
	handler.NewUserHandler(userGroup, userRPC, authMiddleware)
	handler.NewBookHandler(bookGroup, bookRPC, authMiddleware)
	handler.NewAuthorHandler(authorGroup, authorRPC, authMiddleware)
	handler.NewCategoryHandler(categoryGroup, categoryRPC, authMiddleware)
	handler.NewBorrowingRecordHandler(borrowingRecordGroup, userRPC, authMiddleware)

	// Start HTTP server
	fmt.Println("HTTP server is running on port 8080")
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.Data.APPPort)))
}
