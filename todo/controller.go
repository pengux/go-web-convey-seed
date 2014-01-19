package todo

import (
	"github.com/gocraft/web"
)

type Controller struct {
}

func (c *Controller) Create(ctx *Context, rw web.ResponseWriter, req *web.Request) {
	var todo Todo
	App.Ctrl.UnmarshalBody(rw, req, &todo)
	App.Ctrl.Create(Service, rw, &todo, App.Endpoints["todos"])
}

func (c *Controller) Delete(ctx *Context, rw web.ResponseWriter, req *web.Request) {
	App.Ctrl.Delete(Service, rw, req.PathParams["id"])
}

func (c *Controller) Read(ctx *Context, rw web.ResponseWriter, req *web.Request) {
	App.Ctrl.Read(Service, rw, req)
}

func (c *Controller) ReadMany(ctx *Context, rw web.ResponseWriter, req *web.Request) {
	App.Ctrl.ReadMany(Service, rw, req)
}

func (c *Controller) Replace(ctx *Context, rw web.ResponseWriter, req *web.Request) {
	var todoToUpdate Todo
	App.Ctrl.UnmarshalBody(rw, req, &todoToUpdate)
	App.Ctrl.Replace(Service, rw, req.PathParams["id"], &todoToUpdate)
}
