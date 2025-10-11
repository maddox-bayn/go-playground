package strutils

import (
	"reflect"
	"testing"
)

func TestFilterExpensivePrice(t *testing.T) {
	jsonData := []byte(`[
		{"name":"Laptop", "price": 1200},
		{"name": "Phone", "price": 900},
		{"name": "Mouse", "price": 200}
	   ]`)
	got, err := FilterExpensivePrice(jsonData, 500)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	want := []Product{
		{"Laptop", 1200},
		{"Phone", 900},
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("FilterExpensivePrice() %v, want %v", got, want)
	}
}
