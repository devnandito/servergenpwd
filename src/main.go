package main

import (
	"github.com/devnandito/servergenpwd/handlers"
	"github.com/devnandito/servergenpwd/server"
)

func main(){
	http := server.NewServer(":8080")

	http.Handle("GET", "/", handlers.HandleHome)

	// TEMPLATE
	http.File("assets")

	http.Listen()
}