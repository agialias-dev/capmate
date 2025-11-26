package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/agialias-dev/capmate/internal/graph"
)

func main() {
	fmt.Println("Go Graph Tutorial")
	fmt.Println()

	// Load .env files
	// .env.local takes precedence (if present)
	godotenv.Load(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	graphHelper := graph.NewGraphHelper()

	graph.InitializeGraph(graphHelper)

	graph.GreetUser(graphHelper)

	var choice int64 = -1

	for {
		fmt.Println("Please choose one of the following options:")
		fmt.Println("0. Exit")
		fmt.Println("1. Make a Graph call")
		fmt.Println("2. Display Access Token")

		_, err = fmt.Scanf("%d", &choice)
		if err != nil {
			choice = -1
		}

		switch choice {
		case 0:
			// Exit the program
			fmt.Println("Goodbye...")
		case 1:
			// Make a Graph call
			graph.MakeGraphCall(graphHelper)
		case 2:
			// Display Access Token
			graph.DisplayAccessToken(graphHelper)
		default:
			fmt.Println("Invalid choice! Please try again.")
		}

		if choice == 0 {
			break
		}
	}
}
