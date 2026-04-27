package main

import (
	"fmt"
	"os"

	"github.com/administrator/moringa-capstone-go/internal/api"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "joke-cli",
	Short: "A professional CLI tool to fetch jokes",
	Long:  `joke-cli is a fast and beautiful CLI tool built with Go, Cobra, and Pterm to demonstrate advanced Go architecture for the Moringa Capstone.`,
	Run: func(cmd *cobra.Command, args []string) {
		spinner, _ := pterm.DefaultSpinner.Start("Fetching a hilarious joke...")

		joke, err := api.FetchJoke()
		if err != nil {
			spinner.Fail("Failed to fetch a joke!")
			pterm.Error.Println(err)
			os.Exit(1)
		}

		spinner.Success("Joke fetched successfully!")

		pterm.DefaultBasicText.Println()
		pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightMagenta)).Println("Random Joke")
		
		pterm.Info.Println(joke.Setup)
		pterm.Success.Println("... " + joke.Punchline)
		pterm.DefaultBasicText.Println()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
