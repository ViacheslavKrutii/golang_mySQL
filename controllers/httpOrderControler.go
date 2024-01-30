package controllers

import (
	"encoding/json"
	"fmt"
	"golang_mvc_REST_API/db"
	"golang_mvc_REST_API/models"
	"golang_mvc_REST_API/views"
	"io"
	"log"
	"net/http"
)

type OrderController struct {
	db db.DbInterface
}

func NewOrder(newDB db.DbInterface) *OrderController {
	return &OrderController{db: newDB}
}

func (o *OrderController) MakeOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("Make order controller")
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	var newOrder models.Order

	err = json.Unmarshal(body, &newOrder)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	IdOrder := o.db.AddOrder(newOrder)
	w.Write([]byte(fmt.Sprintf("Order added successfully. ID: %d\n", IdOrder)))

}

func (o *OrderController) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete order controller")
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	var newDeleteRequest models.DeleteOrderRequest

	err = json.Unmarshal(body, &newDeleteRequest)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	err = o.db.DeleteOrder(newDeleteRequest)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(fmt.Sprintf("Order deleted successfully. ID: %d\n", newDeleteRequest.IdOrder)))

}

func Example(w http.ResponseWriter, r *http.Request) {

	orderExample := models.Order{OrderBody: map[string]int{"Борщ": 3}, User: models.User{Name: "me"}}

	w.Write(views.JsonOrderExample(&orderExample))
}
