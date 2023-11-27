package views

import (
	"encoding/json"
	"golang_mvc_REST_API/models"
)

func JsonMenu(m *models.Menu) []byte {
	serialized, _ := json.Marshal(m)
	return serialized
}
