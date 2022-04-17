package json

import (
	"encoding/json"
	"testing"
)

type Order struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	TotalPrice float32 `json:"total_price"`
}

func TestJson(t *testing.T) {
	order := Order{
		ID:         "1234",
		Name:       "learn go",
		Quantity:   10,
		TotalPrice: 30,
	}
	data, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}
	t.Logf("%s\n", data)
}

func TestUnMarshal(t *testing.T) {
	str := `{"id":"1234","name":"learn go","quantity":10,"total_price":30}`
	var order Order
	json.Unmarshal([]byte(str), &order)
	t.Logf("%+v\n", order)
	t.Log(order.Name)
}
