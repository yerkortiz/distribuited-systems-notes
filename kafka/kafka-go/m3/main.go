package main

import (
	"log"
	"milestone3/events"
	"milestone3/server"
	"os"
	"time"
)

func init() {
	log.SetOutput(os.Stdout)
	orderReceived := events.OrderReceived{
		EventID:    12312,
		EventState: "Received",
		Timestamp:  time.Now(),
		EventMessage: events.Order{
			OrderID: 1231,
			Total:   1000,
			Customer: events.Customer{
				Email:    "asdas@mail.com",
				FistName: "caca",
				LastName: "seca",
				ShippingAddress: events.ShippingAddress{
					StreetAddress: "casa",
					City:          "ciudad",
					Commune:       "comuna",
					Country:       "pais",
					ZipCode:       "12312312",
				},
			},
			Products: []events.Product{{
				ProductID:   123132,
				Price:       1000,
				ProductName: "inverse suffix tree",
			}},
		},
	}
	events.PublishEvent(orderReceived, 0)
}
func main() {
	server.RunServer()
}
