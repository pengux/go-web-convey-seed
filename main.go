package main

import (
	"github.com/gocraft/web"
	"log"
	"net/http"
)

const (
	Address string = ":9090"
)

var (
	Ctrl      = &Controller{}
	Endpoints = make(map[string]string)
	Routes    = make(map[string]func(app *App))
)

type Context struct {
	Ctrl *Controller
}

func (c *Context) Init(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {

	c.Ctrl = Ctrl
	next(rw, req)
}

type App struct {
	RootRouter *web.Router
}

// Bootstrap an App and initialize routes for endpoints
func Bootstrap() *App {

	app := &App{}
	app.RootRouter = web.New(Context{}).
		Middleware((*Context).Init)

	// Init routes for endpoints
	for _, fn := range Routes {
		fn(app)
	}

	return app
}

// Application's state such as DB connections, global objects etc. should be setup here
func init() {
}

func main() {

	app := Bootstrap()
	log.Print("Listen to port ", Address)
	http.ListenAndServe(Address, app.RootRouter)
}
