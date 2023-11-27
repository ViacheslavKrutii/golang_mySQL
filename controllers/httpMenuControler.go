package controllers

import (
	"golang_mvc_REST_API/models"
	"golang_mvc_REST_API/views"
	"net/http"
)

type MenuController struct {
	Menu *models.Menu
}

func (m *MenuController) AddMenu(Menu *models.Menu) {
	m.Menu = Menu

}

func (m *MenuController) ShowMenuController(w http.ResponseWriter, r *http.Request) {
	w.Write(views.JsonMenu(m.Menu))
}
