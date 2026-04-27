# Getting Started with Go (Golang) – A Beginner’s Guide

## 1. Title & Objective
**Title:** "Prompt-Powered Kickstart: Building a Beginner’s Toolkit for Go"
**Technology Chosen:** Go (Golang)
**Why Chosen:** Go is a powerful, statically typed, compiled language developed at Google. It's known for its simplicity, fast compilation times, and excellent support for concurrency. It's widely used in backend development, microservices, and cloud-native applications. I wanted to learn how to build simple, fast APIs and CLI tools with it.
**End Goal:** Create a minimal working CLI application that makes an HTTP GET request to a public API and parses the JSON response.

## 2. Quick Summary of the Technology
Go (often called Golang) is an open-source programming language supported by Google. It combines the performance of compiled languages like C++ with the ease of use of dynamically typed languages like Python.
**Where is it used?** Go is heavily used in cloud infrastructure, networking tools, and backend web services. Tools like Docker and Kubernetes are built in Go.
**Real-world example:** Building high-performance microservices that require low latency and high concurrency, such as an Uber ride-matching service.

## 3. System Requirements
- **OS:** Linux, Mac, or Windows
- **Tools/Editors required:** VS Code (with the Go extension recommended)
- **Packages:** The Go toolchain (compiler and tools)

## 4. Installation & Setup Instructions

### Step 1: Install Go
Download the installer from the [official Go website](https://go.dev/dl/). 

On Ubuntu/Linux, you can also use:
```bash
sudo apt update
sudo apt install golang-go
```

Verify the installation:
```bash
go version
# Output: go version go1.x.x linux/amd64
```

### Step 2: Initialize a Project
Create a new directory and initialize a Go module:
```bash
mkdir capstone
cd capstone
go mod init github.com/yourusername/capstone-go
```

## 5. Minimal Working Example
This example fetches a random joke from the Official Joke API, parses the JSON, and prints it to the console.

**Code (`main.go`):**
```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Define a struct to map the JSON response
type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func main() {
	fmt.Println("Fetching a random joke for you...")

	// Make an HTTP GET request
	resp, err := http.Get("https://official-joke-api.appspot.com/random_joke")
	if err != nil {
		fmt.Printf("Error fetching joke: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close() // Ensure the connection is closed

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		os.Exit(1)
	}

	// Parse (unmarshal) the JSON into our struct
	var joke Joke
	err = json.Unmarshal(body, &joke)
	if err != nil {
		fmt.Printf("Error parsing joke: %v\n", err)
		os.Exit(1)
	}

	// Print the output
	fmt.Printf("\n%s\n", joke.Setup)
	fmt.Printf("... %s\n\n", joke.Punchline)
}
```

**To run the code:**
```bash
go run main.go
```

**Expected Output:**
```
Fetching a random joke for you...

Why did the invisible man turn down the job offer?
... He couldn't see himself doing it.
```

## 6. AI Prompt Journal

- **Prompt used:** "Give me a step-by-step guide to initialize a new Go project and write a simple script to fetch data from a public JSON API."
- **AI’s response summary:** The AI provided the installation commands, explained the `go mod init` command, and generated a robust script using standard libraries like `net/http` and `encoding/json`.
- **Brief part of the response that addresses the problem:** "To parse JSON in Go, define a `struct` that matches the JSON structure and use `json.Unmarshal()` to map the data."
- **Evaluation:** Very helpful. Go's strict typing makes JSON parsing slightly more verbose than Python or JS, and the AI's explanation of defining structs was crucial for understanding how data mapping works in Go.

- **Prompt used:** "Why do I need `defer resp.Body.Close()` in my Go HTTP request?"
- **AI’s response summary:** The AI explained that HTTP clients in Go leave network connections open to reuse them. Failing to close the response body causes a resource leak.
- **Evaluation:** Excellent. This is a common pitfall for beginners, and the AI highlighted it before it became a problem.

## 7. Common Issues & Fixes
- **Error:** `go: cannot determine module path for source directory`
  - **Fix:** You forgot to initialize the module. Run `go mod init <module-name>` before creating your Go files.
- **Error:** `imported and not used: "fmt"`
  - **Fix:** Go is extremely strict and will not compile if you have unused imports or variables. Remove the import or use the variable.
- **Error:** JSON parsing fails or returns empty fields.
  - **Fix:** In Go, struct fields MUST start with a capital letter to be "exported" (accessible by the `encoding/json` package). Also, ensure your `json:"field_name"` tags exactly match the API's response.

## 8. References
- [Official Go Documentation](https://go.dev/doc/)
- [A Tour of Go](https://go.dev/tour/welcome/1)
- [Go by Example](https://gobyexample.com/)
- [Official Joke API](https://github.com/15Dkatz/official_joke_api)
