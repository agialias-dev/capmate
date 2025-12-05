package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/agialias-dev/capmate/internal/graph"
)

func main() {
	fmt.Println("Welcome to CAPMate!")
	fmt.Println()

	// Load .env file
	godotenv.Load(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	userSession := graph.NewUserSession()
	graph.InitializeGraph(userSession)
	graph.GreetUser(userSession)

	var choice int64 = -1
	for {
		fmt.Println("Please choose one of the following options:")
		fmt.Println("0. Exit")
		fmt.Println("1. Get Conditional Access Policies")
		fmt.Println("2. Get HTTP Conditional Access Policies")

		_, err = fmt.Scanf("%d", &choice)
		if err != nil {
			choice = -1
		}
		switch choice {
		case 0:
			fmt.Println("Goodbye...")
		case 1:
			graph.GetAllCAPs(userSession)
		case 2:
			err = graph.GetHTTPCAPs(userSession)
			if err != nil {
				log.Println(err.Error())
			}
		default:
			fmt.Println("Invalid choice! Please try again.")
		}
		if choice == 0 {
			break
		}
	}
}
