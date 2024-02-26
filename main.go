package main

import (
	"golang_mvc_REST_API/controllers"
	"golang_mvc_REST_API/db"
	"golang_mvc_REST_API/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Print("main start done\n")
	menu1 := &models.Menu{MenuItems: []models.MenuItem{{Name: "Борщ", Price: 10}}}
	log.Print("menu init\n")
	datastore := db.NewPgMenuState()

	orderController := controllers.NewOrder(datastore)
	menuController := controllers.MenuController{}
	menuController.AddMenu(menu1)

	r := mux.NewRouter()

	r.HandleFunc("/menu", menuController.ShowMenu).Methods("GET")
	r.HandleFunc("/order/make", orderController.MakeOrder).Methods("POST")
	r.HandleFunc("/order/delete", orderController.DeleteOrder).Methods("POST")
	r.HandleFunc("/order/example", controllers.Example).Methods("GET")

	http.ListenAndServe("0.0.0.0:8080", r)
}
