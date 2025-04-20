// Package main provides order management functionality.
package main

import "context"

// OrdersService defines the interface for handling order-related business logic operations.
// It serves as the primary service layer interface for order management.
type OrdersService interface {
	// CreateOrder processes a new order creation request.
	// It takes a context.Context for handling timeouts and cancellations.
	// Returns an error if the order creation fails.
	CreateOrder(context.Context) error
}

// OrdersStore defines the interface for order persistence operations.
// It abstracts the underlying storage implementation for order data.
type OrdersStore interface {
	// Create persists a new order in the storage.
	// It takes a context.Context for handling timeouts and cancellations.
	// Returns an error if the storage operation fails.
	Create(context.Context) error
}