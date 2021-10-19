package test

import (
	"milestone3/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HealthCheck)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"service_status_alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestReceiveOrder(t *testing.T) {
	body := strings.NewReader(`{
		"event_id":1231231,
		"event_state":"Received",
		"timestamp":"2018-09-22T12:42:31Z",
		"event_message":{
			"order_id":1993821,
			"total":144,
			"customer":{
				"email":"foo@bar.com",
				"first_name":"Foo",
				"last_name":"Bar",
				"shipping_address": {
					"street_address": "the street 1",
					"city": "city",
					"state": "state",
					"country": "the world",
					"zip_code": "1231221"
				}
			},
			"products": [{
				"product_id":1231,
				"price":12,
				"product_name":"something"
			}]
		}
	}
	`)
	req, err := http.NewRequest("POST", "/order", body)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ReceiveOrder)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"order_received": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
