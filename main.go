package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Hello struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	h := Hello{Message: "hello"}
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(h); err != nil {
		http.Error(w, "the sever encounted an error that could not process your request", http.StatusInternalServerError)
		return
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	// Implement echo functionality here
	defer r.Body.Close()
	w.Header().Set("content-type", "application/json")
	if _, err := io.ReadAll(io.LimitReader(r.Body, 10<<20)); err != nil {
		http.Error(w, "ecountered problem on stream of data", http.StatusInternalServerError)
	}
}

// Implement the FilterExpensiveProducts function and test it in a small Go main (call it directly, not via HTTP).

type Products struct {
	Id    int `json:"id"`
	Name  string `json:"name"`
	Price float64 `json:"price"`
}

func FilterExpensiveProducts(products []Products, threshold float64) []Products {
	var expensiveproduct []Products
	for _, product := range products {
		if product.Price > threshold {
			expensiveproduct = append(expensiveproduct, product)
		}
	}
	return expensiveproduct
}

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	thresholdstr := r.URL.Query().Get("threshold")
	if thresholdstr == "" {
		http.Error(w, "threshold query parameter is required", http.StatusBadRequest)
	}
	threshold, err := strconv.ParseFloat(thresholdstr, 64)
	if err != nil {
		http.Error(w, "invalid threshold value", http.StatusBadRequest)
	}
	result := FilterExpensiveProducts([]Products{
		{Id: 1, Name: "Laptop", Price: 1200.50},
		{Id: 2, Name: "Mouse", Price: 25.00},
		{Id: 3, Name: "Keyboard", Price: 75.00},
		{Id: 4, Name: "Monitor", Price: 300.00},
	}, threshold)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"threshold": threshold,
		"products":  result,
		"count":    len(result),
	})
	if err != nil {
		http.Error(w, "issue with the sver", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/filter", FilterHandler)
	fmt.Println("starting sever on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server %v\n:", err)
	}
	products := []Products{
		{Id: 1, Name: "Laptop", Price: 1200.50},
		{Id: 2, Name: "Mouse", Price: 25.00},
		{Id: 3, Name: "Keyboard", Price: 75.00},
		{Id: 4, Name: "Monitor", Price: 300.00},
	}
	fmt.Println("Testing filterexpensiveproduct function:")
	expensiveProducts := FilterExpensiveProducts(products, 100.00)
	fmt.Println("Expensive Products:", expensiveProducts)
}
