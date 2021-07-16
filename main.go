package main

import (
	"database/sql"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/dsykes16/gofib/cache"
	"github.com/dsykes16/gofib/cache/local_cache"
	"github.com/dsykes16/gofib/cache/pg_cache"
	"github.com/dsykes16/gofib/fibonacci"
	"github.com/dsykes16/gofib/server"

	pb "github.com/dsykes16/gofib/protos"
)

func main() {
	// Load Configuration
	config, err := LoadConfig("./")
	if err != nil {
		log.Fatalf("unable to load config: %s", err)
	}

	// Initialize Memoization Cache
	var c cache.Cache
	switch config.DBDriver {
	case "local":
		c = local_cache.New()
	case "postgres":
		conn := pg_cache.PgConnection{
			Host:     config.DBHost,
			Port:     config.DBPort,
			User:     config.DBUser,
			Password: config.DBPass,
			DbName:   config.DBName,
			SSL:      config.DBSSL,
		}
		db, err := sql.Open("postgres", conn.ConnectionString())
		if err != nil {
			log.Fatalf("could not connect to postgres: %s", err)
		}
		c = pg_cache.New(db)
	}

	// Initialize Fibonacci Backend
	fib := fibonacci.New(c)

	// Initialize Fibonacci Server
	fibServer := server.New(fib)

	// Initialize gRPC Server
	grpcServer := grpc.NewServer()
	pb.RegisterFibonacciServer(grpcServer, fibServer)

	// Start gRPC Server
	lis, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatalf("failed to open port: %s", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc server: %s", err)
	}
}
