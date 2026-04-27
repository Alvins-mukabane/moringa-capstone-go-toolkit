package api

import (
	"testing"
)

func TestFetchJoke(t *testing.T) {
	// A simple test to ensure the external API call succeeds
	// In a real production environment, you would mock the HTTP response.
	joke, err := FetchJoke()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if joke == nil {
		t.Fatal("Expected a joke, got nil")
	}

	if joke.Setup == "" || joke.Punchline == "" {
		t.Errorf("Joke fields should not be empty, got: %+v", joke)
	}
}
