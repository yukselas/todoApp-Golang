package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const VERSION = "0.1.0"
const DEFAULT_PORT = "7890"

var notes []Note

func main() {
	fmt.Printf("todo-app👆 v%s\n", VERSION)
	err := Init()
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		os.Exit(1)
	}
	RunServer()
	fmt.Println("Bye")
}

func Init() error {
	err := godotenv.Load()
	if err != nil {
		return nil
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	notes = append(notes, NewNote("Hello, world!"))
	notes = append(notes, NewNote("Goodbye, world!"))

	err = InitServer()
	if err != nil {
		return err
	}
	return nil
}
