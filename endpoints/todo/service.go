package todo

import (
	"github.com/pengux/go-web-convey-seed/app"
)

type TodoService struct {
	*app.InMemoryDBService
}
