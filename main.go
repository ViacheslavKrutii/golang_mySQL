package main

import (
	"golang_mvc_REST_API/controllers"
	"golang_mvc_REST_API/db"
	"golang_mvc_REST_API/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	menu1 := &models.Menu{MenuItems: []models.MenuItem{{Name: "Борщ", Price: 10}}}
	datastore := db.NewInMemoryState()

	orderController := controllers.NewOrderController(datastore)
	menuController := controllers.MenuController{}
	menuController.AddMenu(menu1)

	r := mux.NewRouter()

	r.HandleFunc("/menu", menuController.ShowMenuController).Methods("GET")
	r.HandleFunc("/order/make", orderController.MakeOrderController).Methods("POST")
	r.HandleFunc("/order/delete", orderController.DeleteOrderController).Methods("POST")
	r.HandleFunc("/order/example", controllers.Example).Methods("GET")

	http.ListenAndServe("localhost:8080", r)
}
