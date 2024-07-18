package main

import (
	"api"
	"server"
)

func main() {
	println("Hello, World!")
	api.Run()
	server.Run()
}
