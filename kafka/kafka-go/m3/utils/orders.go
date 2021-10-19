package utils

import (
	"milestone3/events"
	"time"

	"github.com/asaskevich/govalidator"
)

func ValidateOrder(order events.Order) (bool, error) {
	result, err := govalidator.ValidateStruct(order)
	if err != nil {
		return false, err
	}
	return result, nil
}
func OrderToEvent(order events.Order) events.OrderReceived {
	var event = events.OrderReceived{
		EventID:      1,
		Timestamp:    time.Now(),
		EventState:   "received",
		EventMessage: order,
	}
	return event
}
