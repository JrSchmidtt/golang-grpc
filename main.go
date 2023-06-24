package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/JrSchmidtt/golang-grpc/src/database"
	"github.com/JrSchmidtt/golang-grpc/src/pb"
	"github.com/JrSchmidtt/golang-grpc/src/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a new category
	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	// create grpc server
	grpcServer := grpc.NewServer()

	// register reflection service on grpc server for debugging purposes.
	// this is optional but can be useful to inspect the service using grpc_cli
	reflection.Register(grpcServer)

	// register category service into grpc server
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	// run grpc server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	// initialize list of grpc servers
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
	fmt.Println("Server running on port :50051")
}