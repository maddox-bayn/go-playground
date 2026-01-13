package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	if _, err := io.ReadAll(io.LimitReader(r.Body, 10 << 20)); err != nil {
		http.Error(w, "ecountered problem on stream of data", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/echo", echo)
	fmt.Println("starting sever on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to start server %v\n:", err)
	}
}
