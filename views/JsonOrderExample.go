package views

import (
	"encoding/json"
	"golang_mvc_REST_API/models"
)

func JsonOrderExample(o *models.Order) []byte {
	serialized, _ := json.Marshal(o)
	return serialized
}
