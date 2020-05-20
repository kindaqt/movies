package main

import (
	"fmt"
	"hello/server"
)

func main() {
	fmt.Println("hello, world")

	router := server.Router()
	router.Run()
}
