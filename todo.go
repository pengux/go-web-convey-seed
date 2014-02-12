package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/pengux/web"
)

type TodoService struct {
	*InMemoryDBService
}

type TodoContext struct {
	*Context
	TodoService *TodoService
}

type Todo struct {
	Id        string
	CreatedAt time.Time
	Title     string `json:"title"`
	Done      bool   `json:"done"`
}

func (d *Todo) GetId() string            { return d.Id }
func (d *Todo) SetId(id string)          { d.Id = id }
func (d *Todo) GetCreatedAt() time.Time  { return d.CreatedAt }
func (d *Todo) SetCreatedAt(t time.Time) { d.CreatedAt = t }

type Todos []*Todo

func (todos *Todos) Append(object interface{}) {
	todo := object.(*Todo)
	*todos = append(*todos, todo)
}

func (c *TodoContext) Init(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	// Init DB service
	dbService := &InMemoryDBService{}
	dbService.Init()

	c.TodoService = &TodoService{dbService}

	next(rw, req)
}

func (c *TodoContext) Create(rw web.ResponseWriter, req *web.Request) {
	todo := &Todo{}
	c.UnmarshalBody(rw, req, &todo)
	todo.Id = ""
	todo.CreatedAt = time.Now()

	id, err := c.TodoService.Create(todo)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Panicln(err)
	}

	rw.Header().Set("Location", "/todos/"+id)
	rw.WriteHeader(http.StatusCreated)
}

func (c *TodoContext) Delete(rw web.ResponseWriter, req *web.Request) {
	err := c.TodoService.Delete(req.PathParams["id"])
	if err != nil {
		switch err := err.(type) {
		case RowNotFoundError:
			rw.WriteHeader(http.StatusNotFound)
		default:
			rw.WriteHeader(http.StatusInternalServerError)
			log.Panicln(err)
		}
	}

	rw.WriteHeader(http.StatusOK)
}

func (c *TodoContext) Read(rw web.ResponseWriter, req *web.Request) {
	todo := &Todo{}
	result, err := c.TodoService.Read(req.PathParams["id"], todo)
	if err != nil {
		switch err.(type) {
		case RowNotFoundError:
			rw.WriteHeader(http.StatusNotFound)
		default:
			rw.WriteHeader(http.StatusInternalServerError)
			log.Panicln(err)
		}
	}

	resp, err := json.MarshalIndent(&result, "", "    ")
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Panicln(err)
	}

	rw.Write(resp)
}

func (c *TodoContext) ReadMany(rw web.ResponseWriter, req *web.Request) {
	todos := &Todos{}
	err := c.TodoService.ReadMany(todos)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Panicln(err)
	}

	resp, err := json.MarshalIndent(todos, "", "    ")
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Panicln(err)
	}

	rw.Write(resp)
}

func (c *TodoContext) Replace(rw web.ResponseWriter, req *web.Request) {
	todoToUpdate := &Todo{}
	c.UnmarshalBody(rw, req, &todoToUpdate)
	err := c.TodoService.Update(req.PathParams["id"], todoToUpdate)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Panicln(err)
	}

	rw.WriteHeader(http.StatusOK)
}
