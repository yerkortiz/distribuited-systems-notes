package handlers

import (
	"encoding/json"
	"io"
	"milestone3/events"
	"milestone3/utils"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"service_status_alive": true}`)
}

func ReceiveOrder(w http.ResponseWriter, r *http.Request) {
	var order events.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	flag, err := utils.ValidateOrder(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !flag {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	event := utils.OrderToEvent(order)
	events.PublishEvent(event, 0)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"order_received": true}`)
}
