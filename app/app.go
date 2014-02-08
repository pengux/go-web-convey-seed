package app

import (
	"fmt"

	"github.com/gocraft/web"
)

var (
	Bootstraps = make(map[string]func(app *Application, context interface{}))
)

type Context struct{}

type Application struct {
	Prefix     string
	RootRouter *web.Router
	Endpoints  map[string]string
}

// Create a new App and initialize routes for endpoints
func New(prefix string, contexts map[string]interface{}) *Application {
	// Check that there is a context for "app"
	if _, ok := contexts["app"]; !ok {
		panic(fmt.Sprintf("You must provide a context for 'app'"))
	}

	newApp := &Application{
		prefix,
		web.NewWithPrefix(contexts["app"], prefix),
		make(map[string]string),
	}

	// Run bootstrap functions for endpoints
	for endpoint, fn := range Bootstraps {
		if endpointCtx, ok := contexts[endpoint]; ok {
			fn(newApp, endpointCtx)
		} else {
			fn(newApp, contexts["app"])
		}
	}

	return newApp
}

// Get endpoint with prefix
func (a *Application) GetEndpointWithPrefix(key string) string {
	endpoint, ok := a.Endpoints[key]
	if !ok {
		panic(fmt.Sprintf("Endpoint %s was not found", key))
	}

	return a.Prefix + endpoint
}
