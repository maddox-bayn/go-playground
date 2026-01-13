package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/maddox-bayn/go-playground/strutils"
)

type helloResp struct {
	Name      string `json:"name"`
	RuneCount int    `json:"rune_count"`
	FirstRune string `json:"first_rune,omitempty"`
	LastRune  string `json:"last_rune,omitempty"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}

	resp := helloResp{
		Name:      name,
		RuneCount: strutils.RuneCount(name),
	}
	if name != "" {
		resp.FirstRune = fmt.Sprintf("%q", strutils.FirstRune(name))
		resp.LastRune = fmt.Sprintf("%q", strutils.LastRune(name))
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func main() {
	// If no CLI arg provided, start HTTP server.
	if len(os.Args) < 2 {
		http.HandleFunc("/hello", helloHandler)
		addr := ":8080"
		log.Printf("Starting server on %s", addr)
		log.Fatal(http.ListenAndServe(addr, nil))
		return
	}

	// CLI mode: show string utilities for the provided word.
	word := os.Args[1]
	fmt.Printf("Word: %v\n", word)
	fmt.Printf("RuneCount: %d\n", strutils.RuneCount(word))
	fmt.Printf("FirstRune: %q\n", strutils.FirstRune(word))
	fmt.Printf("LastRune: %q\n", strutils.LastRune(word))
}
