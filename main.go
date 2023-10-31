package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.huddl.one/rce/src/api"
	"go.huddl.one/rce/src/docker"
)

func DisplayHelpMenu() {
	fmt.Println("Usage: ./rce <command>")
	fmt.Println("Commands:")
	fmt.Println("  serve - start the server")
	fmt.Println("  help - show this help")
}

func main() {

	godotenv.Load()

	PORT := os.Getenv("PORT")

	if PORT == "" {
		log.Fatal("$PORT must be set in .env file. Make a copy of .env.example and rename it to .env and set the port number")
	}

	switch os.Args[1] {
	case "serve":
		api.Serve(PORT)
	case "pull-images":
		docker.PullImages()
	case "help":
		DisplayHelpMenu()
	default:
		DisplayHelpMenu()
	}

}
