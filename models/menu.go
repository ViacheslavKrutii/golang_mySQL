package models

type Hryvna float32

type MenuItem struct {
	Name  string `json:"dish"`
	Price Hryvna `json:"prise"`
}

type Menu struct {
	MenuItems []MenuItem `json:"positions"`
}
