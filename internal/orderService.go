package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/dawitel/E-Commerce_app_order-service/protogen/golang/orders"
)

// The orderService implements the order_service interface created from gRPC
// UnimplementedOrdersServer must be embedded to have forwarded compatible implementations.

type OrderService struct {
	db *DB
	orders.UnimplementedOrdersServer
}

func NewOrderService(db *DB) OrderService { // create a new OrderService
	return OrderService{db: db}
}

// AddOrder implements the AddOrder method of the grpc OrdersServer interface to add a new order
func (o *OrderService) AddOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	fmt.Println("Recieved a new Order with orderID %v", req.Order.OrderId)
	err := o.db.AddOrder(req.GetOrder())
	
	return &orders.Empty, err
}

// GetOrder implements the GetOrder method of the grpc OrdersServer interface to fetch an order for a given orderID
func (o *OrderService) GetOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.PayloadWithSingleOrder, error) {
	log.Printf("Recived get order request")

	order := o.db.GetOrderByID((req.GetOrderId()))

	if order == nil {
		return nil, fmt.Errorf("order not found for the orderId: %d", req.GetOrderId())
	}

	return &orders.PayloadWithSingleOrder{Order: order}, nil
}

// UpdateOrder imlements the UpdateOrder method of the grpc OrderServer interface to update the order
func (o *OrderService) UpdateOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Printf("Received an update order request")

	o.UpdateOrder(req.GetOrder())
	
	return &orders.Empty{}, nil
}

// RemoveOrder implements the RemoveOrder method of the grpc OrdersServer interface to remove an order
func (o *OrderService) RemoveOrder(_ context.Context, req *orders.PayloadWithOrderID) (*orders.Empty, error) {
	log.Printf("Received a remove order request")
	
	o.db.RemoveOrder(req.GetOrderId())
	
	return &orders.Empty{}, nil
}
