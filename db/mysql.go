package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"golang_mvc_REST_API/models"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlMenuState struct {
	db *sql.DB
}

func NewMysqlConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/menu")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

func NewMysqlMenuState() *MysqlMenuState {
	return &MysqlMenuState{
		NewMysqlConnection(),
	}
}

func (m *MysqlMenuState) AddOrder(newOrder models.Order) {
	userBody, err := json.Marshal(newOrder.OrderBody)
	if err != nil {
		log.Println(err)
		return
	}
	result, err := m.db.Exec("INSERT INTO orders (Username, order_body,) VALUES (?, ?)", newOrder.User.Name, userBody)
	if err != nil {
		log.Println(err)
		return
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Order added successfully. ID: %d\n", lastInsertID)

}

func (m *MysqlMenuState) DeleteOrder(newDeleteRequest models.DeleteOrderRequest) error {
	result, err := m.db.Exec("DELETE FROM orders WHERE orderID = ?", newDeleteRequest.IdOrder)
	if err != nil {
		log.Println(err)
		return err
	}

	// Перевірка кількості видалених рядків (optional)
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Println("No rows were deleted. Order with specified ID not found.")
		return errors.New("order not found")
	}

	log.Printf("Order with ID %d deleted successfully.\n", newDeleteRequest.IdOrder)
	return nil
}
