package main

import (
	"api-rest-go/database"
	"api-rest-go/routes"
	"fmt"
)

func main() {
	database.Connect()
	fmt.Println("the fun started")
	routes.Init()
}
