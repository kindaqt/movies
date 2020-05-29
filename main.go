package main

import (
	"fmt"

	server "github.com/kindaqt/movies/api/server"
)

func main() {
	fmt.Println("hello, world")

	router := server.Router()
	router.Run()
}
