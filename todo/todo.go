package todo

import (
	"github.com/pengux/go-web-convey-seed/app"
)

const (
	ENDPOINT     = "todos"
	ENDPOINT_URL = "/" + ENDPOINT
)

var (
	Service *TodoService
)

type Context struct {
	*app.Context
}

type Todo struct {
	app.Entity
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type Todos []*Todo

func (todos *Todos) Append(object interface{}) {
	todo := object.(*Todo)
	*todos = append(*todos, todo)
}

func init() {
	app.Bootstraps[ENDPOINT] = func(currentApp *app.Application, context interface{}) {

		currentApp.Endpoints[ENDPOINT] = ENDPOINT_URL

		dbService := &app.InMemoryDBService{}
		dbService.Init()

		Ctrl := &Controller{
			&app.Controller{},
			currentApp,
			&TodoService{
				dbService,
			},
		}

		currentApp.RootRouter.Subrouter(context, currentApp.Endpoints["todos"]).
			Get("/", Ctrl.ReadMany).
			Get("/:id", Ctrl.Read).
			Post("/", Ctrl.Create).
			Put("/:id", Ctrl.Replace).
			Delete("/:id", Ctrl.Delete)
	}
}
