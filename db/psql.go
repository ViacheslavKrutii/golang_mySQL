package db

import (
	"context"
	"encoding/json"
	"fmt"
	"golang_mvc_REST_API/models"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type PgMenuState struct {
	db *pgx.Conn
}

func NewPgConnection() *pgx.Conn {
	db, err := pgx.Connect(context.Background(), "postgres://root:root@postgres:5432/mydb")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	}

	return db
}

func NewPgMenuState() *PgMenuState {
	return &PgMenuState{
		NewPgConnection(),
	}
}

func (p *PgMenuState) AddOrder(newOrder models.Order) (IdOrder int) {
	userBody, err := json.Marshal(newOrder.OrderBody)
	if err != nil {
		log.Println(err)
		return
	}

	var lastInsertID int
	err = p.db.QueryRow(context.Background(), "INSERT INTO orders (Username, order_body) VALUES ($1, $2) RETURNING orderID", newOrder.User.Name, userBody).Scan(&lastInsertID)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Order added successfully. ID: %d\n", lastInsertID)
	return lastInsertID
}

func (p *PgMenuState) DeleteOrder(newDeleteRequest models.DeleteOrderRequest) error {
	_, err := p.db.Query(context.Background(), "DELETE FROM orders WHERE orderID = $1", newDeleteRequest.IdOrder)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Order with ID %d deleted successfully.\n", newDeleteRequest.IdOrder)
	return nil

}
