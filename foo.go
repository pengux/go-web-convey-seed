package main

import (
	"github.com/gocraft/web"
)

type FooContext struct {
	*Context
	Service *FooService
}

func (c *FooContext) Init(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {

	c.Service = &FooService{
		&DBService{
			"foos",
		},
	}
	next(rw, req)
}

type Foo struct {
	Entity
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func init() {
	Endpoints["foos"] = "/foos"
	Routes["foos"] = func(app *App) {

		fooCtrl := &FooController{}

		app.RootRouter.Subrouter(FooContext{}, Endpoints["foos"]).
			Middleware((*FooContext).Init).
			Get("/", fooCtrl.ReadMany).
			Get("/:id", fooCtrl.Read).
			Post("/", fooCtrl.Create).
			Put("/:id", fooCtrl.Replace).
			Delete("/:id", fooCtrl.Delete)
	}
}
