package main

import (
	"fmt"
	"log"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/dawitel/grpc-go/protogen/golang/orders"
	"github.com/dawitel/grpc-go/protogen/golang/product"
)

func main() {
	// Initialize a new order with the Order type from the protogen
	orderItems := orders.Order{
		OrderId: 10,
		CustomerId: 11,
		IsActive: true,
		OrderDate: &date.Date{Year: 2021, Month: 1, Day: 1},
		Products: []*product.Product{
			{ProductId: 1, ProductName: "optimetrics.ai", ProductType: product.ProductType_FOOD},
		},
	}
	// marshal the orderItems into json
	bytes, err := protojson.Marshal(&orderItems)
	if err != nil {
		log.Fatal("DESERIALIZATION_ERROR: ", err)
	}

	fmt.Println((string(bytes)))

}