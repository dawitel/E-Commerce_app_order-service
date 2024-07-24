package internal

import (
	"fmt"

	"github.com/dawitel/grpc-go/proto/orders"
)

type DB struct {
	collection []*orders.Order
}

// NewDB creates a mock in-memory Database instance
func NewDb() *DB {
	return &DB{
		collection: make([]*orders.Order, 0),
	}
}

// AddOrder adds a new order in the mock DB 
// and returns an error if duplicate orderId exists

func (db *DB) AddOrder(order *orders.Order) error {
	for _, o := range db.collection {
		if o.OrderId == order.OrderId { // check for duplicate orders
			return fmt.Errorf("Duplicate ORDER_ID: %v", order.GetOrderId())
		}
	}
	db.collection = append(db.collection, order)
	return nil
}

// GetOrderById returns an order using the order_id
func (db *DB) GetOrderByID(orderId uint64) *orders.Order {
	for _, o := range db.collection {
		if o.OrderId == orderId {
			return o
		}
	}
	return nil
}

// GetOrdersByIDs returns all orders for the given order ids
func (db *DB) GetOrdersByIDs(orderIDs []uint64) []*orders.Order {
	filtered := make([]*orders.Order, 0)

	for _, idx := range orderIDs {
		for _, order := range db.collection {
			if order.OrderId == idx {
				filtered = append(filtered, order)
				break
			}
		}
	}

	return filtered
}

// UpdateOrder updates an order 
func (db *DB) UpdateOrder(order *orders.Order) {
	for i, ord := range db.collection {
		if ord.OrderId == order.OrderId {
			db.collection[i] = order
			return
		}
	}
}

// RemoveOrder takes in an order Id and removes it from the DB
func (db *DB) RemoveOrder(orderId uint64) {
	filtered :=  make([]*orders.Order, 0, len(db.collection)-1)
	for i := range db.collection {
		if db.collection[i].OrderId != orderId {
			filtered = append(filtered, db.collection[i])
		}
	}
	db.collection = filtered
}
