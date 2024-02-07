package db

import (
	"context"
	"fmt"
	"golang_mvc_REST_API/models"
	"os"

	"github.com/jackc/pgx/v5"
)

type PgMenuState struct {
	db *pgx.Conn
}

func NewPgConnection() *pgx.Conn {
	db, err := pgx.Connect(context.Background(), "user=root password=root host=localhost port=5432 dbname=mydb sslmode=verify-ca")

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

}
