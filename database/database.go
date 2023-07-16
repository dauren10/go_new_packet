package database

import "fmt"

func Connect() {
	fmt.Println("Connecting to the database...")
}

func Disconnect() {
	fmt.Println("Close connecting to the database...")
}

type User struct {
	Name  string
	Email string
}
