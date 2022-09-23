package models

type PersonResponse struct {
	Status struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	} `json:"status"`
	Person []PersonModel `json:"result"`
}

type PersonModel struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	UUID      string `json:"uuid"`
}

type PersonOrders struct {
	Person PersonModel `json:"person"`
	Orders []Order     `json:"orders"`
}
