package models

type Order struct {
	OrderBody map[string]int `json:"order"`
	User      User           `json:"user"`
}

type DeleteOrderRequest struct {
	User    User `json:"user"`
	IdOrder int`json:"idOrder"`
}
