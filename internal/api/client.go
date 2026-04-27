package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

// FetchJoke fetches a random joke from the official joke API.
func FetchJoke() (*Joke, error) {
	resp, err := http.Get("https://official-joke-api.appspot.com/random_joke")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch joke: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var joke Joke
	if err := json.Unmarshal(body, &joke); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return &joke, nil
}
