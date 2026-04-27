package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pterm/pterm"
)

type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

var jokes = []Joke{
	{Setup: "Why do programmers prefer dark mode?", Punchline: "Because light attracts bugs!"},
	{Setup: "How many programmers does it take to change a light bulb?", Punchline: "None. It's a hardware problem."},
	{Setup: "What's a programmer's favorite hangout place?", Punchline: "Foo Bar."},
}

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Just return the first joke as a simple API example
	joke := jokes[0]

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(joke); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main() {
	port := "8080"
	http.HandleFunc("/api/joke", jokeHandler)

	pterm.Success.Printf("Starting Joke API Server on http://localhost:%s\n", port)
	pterm.Info.Println("Test it: curl http://localhost:8080/api/joke")

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
