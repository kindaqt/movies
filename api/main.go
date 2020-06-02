package main

import (
	"fmt"

	server "github.com/kindaqt/movies/api/server"
)

func main() {
	fmt.Println("Starting Movies API")

	router := server.Router()
	router.Run()
}
