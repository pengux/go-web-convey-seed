package main

import (
	"github.com/gocraft/web"
)

type FooController struct {
}

func (c *FooController) Create(ctx *FooContext, rw web.ResponseWriter, req *web.Request) {

	var foo Foo
	ctx.Ctrl.UnmarshalBody(rw, req, &foo)
	ctx.Ctrl.Create(ctx.Service, rw, &foo, Endpoints["foos"])
}

func (c *FooController) Delete(ctx *FooContext, rw web.ResponseWriter, req *web.Request) {

	ctx.Ctrl.Delete(ctx.Service, rw, req.PathParams["id"])
}

func (c *FooController) Read(ctx *FooContext, rw web.ResponseWriter, req *web.Request) {

	var foo Foo
	ctx.Ctrl.Read(ctx.Service, rw, req.PathParams["id"], &foo)
}

func (c *FooController) ReadMany(ctx *FooContext, rw web.ResponseWriter, req *web.Request) {

	var foos []Foo
	ctx.Ctrl.ReadMany(ctx.Service, rw, &foos)
}

func (c *FooController) Replace(ctx *FooContext, rw web.ResponseWriter, req *web.Request) {

	var fooToUpdate Foo
	ctx.Ctrl.UnmarshalBody(rw, req, &fooToUpdate)
	ctx.Ctrl.Replace(ctx.Service, rw, req.PathParams["id"], &fooToUpdate)
}
