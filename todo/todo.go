package todo

import (
	"github.com/pengux/go-web-convey-seed/app"
)

var (
	Service *TodoService
	App     *app.Application
)

type Context struct {
	*app.Context
}

type Todo struct {
	app.Entity
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func init() {
	app.Bootstraps["todos"] = func(currentApp *app.Application) {

		currentApp.Endpoints["todos"] = "/todos"
		App = currentApp

		dbService := &app.InMemoryDBService{}
		dbService.Init()
		Service = &TodoService{
			dbService,
		}

		Ctrl := &Controller{}

		currentApp.RootRouter.Subrouter(Context{}, currentApp.Endpoints["todos"]).
			Get("/", Ctrl.ReadMany).
			Get("/:id", Ctrl.Read).
			Post("/", Ctrl.Create).
			Put("/:id", Ctrl.Replace).
			Delete("/:id", Ctrl.Delete)
	}
}
