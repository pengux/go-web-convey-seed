package main

import (
	"log"
	"net/http"

	"github.com/pengux/go-web-convey-seed/app"
	"github.com/pengux/go-web-convey-seed/todo"
)

const (
	Address string = ":9090"
)

func main() {
	mainContext := &app.Context{}
	app := app.New("", map[string]interface{}{
		"app": mainContext,
		"todos": &todo.Context{
			mainContext,
		},
	})
	log.Print("Listen to port ", Address)
	http.ListenAndServe(Address, app.RootRouter)
}
