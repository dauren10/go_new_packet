package main

import (
	"fmt"

	"dyns.com/myapp/database"
)

func main() {
	fmt.Println("Hell")
	database.Connect()
	user := database.User{Name: "John", Email: "john@example.com"}
	fmt.Println(user)
	database.Disconnect()

}
