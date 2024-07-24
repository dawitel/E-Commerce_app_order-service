package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dawitel[grpc-go/protogen/golang/orders"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/genproto/googleapis/cloud/bigquery/connection/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var orderServiceAddr string

func main() {
	// setup a connection to the order service
	fmt.Println("Connecting to order service via: ", orderServiceAddr)
	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to order service: %v", err)
	}

	defer conn.close()
	// register the grpc server and chec for the health
	mux := runtime.NewServerMux()
	if err = orders.RegisterOrdersHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register the order server: %v", err)
	}

	// serve the API gateway
	addr := "0.0.0.0:8080"
	fmt.Println("API gateway server is running on: %v", addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("gateway server closed suddenly: %v", err)
	}
}