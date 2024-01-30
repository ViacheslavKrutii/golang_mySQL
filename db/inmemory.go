package db

import (
	"errors"
	"golang_mvc_REST_API/models"
	"log"
)

type InMemoryState struct {
	orders       map[int]models.Order
	lastInsertID int
}

func NewInMemoryState() *InMemoryState {
	log.Print("memory state done\n")
	return &InMemoryState{orders: make(map[int]models.Order)}
}

func (i *InMemoryState) AddOrder(newOrder models.Order) (IdOrder int) {
	IdOrder = i.lastInsertID + 1
	i.orders[IdOrder] = newOrder
	i.lastInsertID = IdOrder
	log.Printf("Order added successfully. ID: %d\n", IdOrder)
	return IdOrder
}

func (i *InMemoryState) DeleteOrder(newDeleteRequest models.DeleteOrderRequest) error {
	_, ok := i.orders[newDeleteRequest.IdOrder]
	if !ok {
		err := errors.New("invalid order index")
		return err
	}

	delete(i.orders, newDeleteRequest.IdOrder)
	log.Printf("Order deleted successfully. ID: %d\n", newDeleteRequest.IdOrder)
	return nil
}
