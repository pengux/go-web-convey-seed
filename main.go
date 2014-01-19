package main

import (
	"github.com/pengux/go-web-convey-seed/app"
	"log"
	"net/http"
)

const (
	Address string = ":9090"
)

func main() {

	App := app.New()
	log.Print("Listen to port ", Address)
	http.ListenAndServe(Address, App.RootRouter)
}
