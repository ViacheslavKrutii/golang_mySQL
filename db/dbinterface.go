package db

import (
	"golang_mvc_REST_API/models"
)

type DbInterface interface {
	AddOrder(models.Order)(IdOrder int)
	DeleteOrder(models.DeleteOrderRequest) error
}
