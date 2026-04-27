# Moringa AI Capstone - Go Professional Architecture

A professional-grade Go (Golang) project demonstrating advanced architecture, standard project layouts, CLI frameworks (Cobra/Pterm), and an internal REST API server.

## Features
- **Joke CLI:** A beautiful, fast command-line interface that fetches random jokes from a public API.
- **Joke API:** A local REST API server that serves its own jokes.
- **Unit Testing:** Basic tests verifying API client connectivity.
- **Automation:** A `Makefile` to simplify builds and execution.

## Running the Code
1. Install Go from [golang.org](https://go.dev/).
2. Clone this repository.
3. Run `make tidy` to install dependencies.
4. Run the CLI: `make run-cli`
5. Run the local API Server: `make run-api`
6. Run unit tests: `make test`
