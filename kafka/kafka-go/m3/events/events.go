package events

import "time"

type OrderReceived struct {
	EventID      int       `json:"event_id"`
	EventState   string    `json:"event_state"`
	Timestamp    time.Time `json:"timestamp"`
	EventMessage Order     `json:"event_message"`
}
type Order struct {
	OrderID  int       `json:"order_id"`
	Total    int       `json:"total"`
	Customer Customer  `json:"customer"`
	Products []Product `json:"products"`
}
type Customer struct {
	Email           string          `json:"email" valid:"email"`
	FistName        string          `json:"first_name"`
	LastName        string          `json:"last_name"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}
type ShippingAddress struct {
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	Commune       string `json:"commune"`
	Country       string `json:"country"`
	ZipCode       string `json:"zip_code"`
}
type Product struct {
	ProductID   int    `json:"product_id"`
	Price       int    `json:"price"`
	ProductName string `json:"product_name"`
}
