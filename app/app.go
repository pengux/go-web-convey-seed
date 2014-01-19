package app

import (
	"fmt"

	"github.com/gocraft/web"
)

var (
	Bootstraps = make(map[string]func(app *Application))
)

type Context struct{}

type Application struct {
	RootRouter *web.Router
	Ctrl       *Controller
	Endpoints  map[string]string
	Prefix     string
}

// Create a new App and initialize routes for endpoints
func New() *Application {
	app := &Application{}
	app.Prefix = "/api"
	app.RootRouter = web.NewWithPrefix(Context{}, app.Prefix)
	app.Ctrl = &Controller{}
	app.Endpoints = make(map[string]string)

	// Run bootstrap functions for endpoints
	for _, fn := range Bootstraps {
		fn(app)
	}

	return app
}

// Get endpoint with prefix
func (app *Application) GetEndpointWithPrefix(key string) string {
	endpoint, ok := app.Endpoints[key]
	if !ok {
		panic(fmt.Sprintf("No endpoint %s found", key))
	}

	return app.Prefix + endpoint
}
