package todo

import (
	"time"
	"github.com/pengux/go-web-convey-seed/app"

	"github.com/gocraft/web"
)

type Controller struct {
	*app.Controller
	App     *app.Application
	Service *TodoService
}

func (c *Controller) Create(ctx *Context, rw web.ResponseWriter, req *web.Request) {
	var todo Todo
	c.Controller.UnmarshalBody(rw, req, &todo)
	todo.SetId("")
	todo.SetCreatedAt(time.Time{})
	c.Controller.Create(c.Service, rw, &todo, c.App.GetEndpointWithPrefix(ENDPOINT))
}

func (c *Controller) Delete(ctx *Context, rw web.ResponseWriter, req *web.Request) {
	c.Controller.Delete(c.Service, rw, req.PathParams["id"])
}

func (c *Controller) Read(ctx *Context, rw web.ResponseWriter, req *web.Request) {
	var todo Todo
	c.Controller.Read(c.Service, rw, req, &todo)
}

func (c *Controller) ReadMany(ctx *Context, rw web.ResponseWriter, req *web.Request) {
	var todos Todos
	c.Controller.ReadMany(c.Service, rw, req, &todos)
}

func (c *Controller) Replace(ctx *Context, rw web.ResponseWriter, req *web.Request) {
	var todoToUpdate Todo
	c.Controller.UnmarshalBody(rw, req, &todoToUpdate)
	c.Controller.Replace(c.Service, rw, req.PathParams["id"], &todoToUpdate)
}
