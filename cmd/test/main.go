package main

import (
	"context"
	"log"
	"net"
	"time"

	"rice-wine-shop/core/adapters/interfaces"
	"rice-wine-shop/core/generator"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen on port 5000: %v", err)
	}

	grpcServer := grpc.NewServer()

	generator.RegisterOrderServiceServer(grpcServer, interfaces.NewOrderServerService())

	go func() {
		log.Println("gRPC server is running on port 5000")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	orderClient := interfaces.NewOrderServerService()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &generator.CreateOrderRequest{Name: "máy tính", Price: 77777}
	_, err = orderClient.CreateOrder(ctx, req)
	if err != nil {
		log.Printf("CreateOrder failed: %v", err)
	} else {
		log.Printf("CreateOrder response")
	}

	select {}
}
