package db

import (
	"errors"
	"golang_mvc_REST_API/models"
	"slices"
)

type InMemoryState struct {
	orders map[models.User][]models.Order
}

func NewInMemoryState() *InMemoryState {
	return &InMemoryState{orders: make(map[models.User][]models.Order)}
}

func (i *InMemoryState) AddOrder(newOrder models.Order) {

	ordersSlice := append([]models.Order(nil), i.orders[newOrder.User]...)
	ordersSlice = append(ordersSlice, newOrder)
	i.orders[newOrder.User] = ordersSlice
}

func (i *InMemoryState) DeleteOrder(newDeleteRequest models.DeleteOrderRequest) error {

	ordersSlice, ok := i.orders[newDeleteRequest.User]
	if !ok || newDeleteRequest.IdOrder >= len(ordersSlice) {
		err := errors.New("invalid order index")
		return err
	}

	i.orders[newDeleteRequest.User] = slices.Delete(ordersSlice, newDeleteRequest.IdOrder, newDeleteRequest.IdOrder)
	return nil
}
