package main

import (
	"log"
	"net"
	
	"github.com/dawitel/grpc-go/internal"
	"github.com/dawitel/grpc-go/protogen/golang/orders"
	
	"google.golang.org/grpc"
)

func main() {
	// create a TCP listener
	const addr = "0.0.0.0:50501"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen in: %v", addr)
	} 	

	server := grpc.NewServer()
	
	// create a new orderservice with a new DB instance
	db := internal.NewDb()
	orderService := internal.NewOrderService(db)

	orders.RegisterOrdersServer(server, &orderService)

	log.Printf("Server is running on %v", listener.Addr())
	if err = server.Serve(listener); err!= nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}