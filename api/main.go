package main

import (
	"fmt"

	server "github.com/kindaqt/movies/api/server"
)

const (
	port string = ":8080"
)

func main() {
	fmt.Println("Starting Movies API")

	router := server.Router()
	router.Run(port)
}
