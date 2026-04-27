package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func main() {
	fmt.Println("Fetching a random joke for you...")

	resp, err := http.Get("https://official-joke-api.appspot.com/random_joke")
	if err != nil {
		fmt.Printf("Error fetching joke: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		os.Exit(1)
	}

	var joke Joke
	err = json.Unmarshal(body, &joke)
	if err != nil {
		fmt.Printf("Error parsing joke: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n%s\n", joke.Setup)
	fmt.Printf("... %s\n\n", joke.Punchline)
}
