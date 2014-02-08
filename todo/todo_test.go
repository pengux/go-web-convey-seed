package todo_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pengux/go-web-convey-seed/app"
	"github.com/pengux/go-web-convey-seed/todo"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTodo(t *testing.T) {

	mainContext := app.Context{}
	todoTestApp := app.New("", map[string]interface{}{
		"app": mainContext,
		"todos": todo.Context{
			&mainContext,
		},
	})

	Convey("Given the todos endpoint", t, func() {
		var testTodo todo.Todo

		Convey("When adding a todo to the collection", func() {
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("POST", todoTestApp.GetEndpointWithPrefix("todos"), strings.NewReader(`{
				"title": "Master Go",
				"done": false
			}`))

			todoTestApp.RootRouter.ServeHTTP(recorder, request)

			Convey("It should returns a 201 Created status", func() {
				So(recorder.Code, ShouldEqual, 201)
			})

		})

		Convey("When listing all todos from the collection", func() {
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", todoTestApp.GetEndpointWithPrefix("todos"), nil)

			todoTestApp.RootRouter.ServeHTTP(recorder, request)

			Convey("It should contains the 'Testtodo' object(s)", func() {
				var todos []todo.Todo
				json.Unmarshal(recorder.Body.Bytes(), &todos)
				testTodo = todos[0]

				So(testTodo.Title, ShouldEqual, "Master Go")
			})
		})

		Convey("When updating the 'Testtodo'", func() {
			testTodo.Done = true
			data, _ := json.Marshal(testTodo)

			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("PUT", todoTestApp.GetEndpointWithPrefix("todos")+"/"+testTodo.Id, strings.NewReader(string(data)))

			todoTestApp.RootRouter.ServeHTTP(recorder, request)

			Convey("the data should persist", func() {
				recorder := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", todoTestApp.GetEndpointWithPrefix("todos")+"/"+testTodo.Id, nil)

				todoTestApp.RootRouter.ServeHTTP(recorder, request)
				var testTodoUpdated todo.Todo
				json.Unmarshal(recorder.Body.Bytes(), &testTodoUpdated)

				So(testTodoUpdated.Done, ShouldEqual, true)
			})
		})

		Convey("When deleting the 'Testtodo'", func() {
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("DELETE", todoTestApp.GetEndpointWithPrefix("todos")+"/"+testTodo.Id, nil)

			todoTestApp.RootRouter.ServeHTTP(recorder, request)

			Convey("it should returns 200 Status", func() {
				So(recorder.Code, ShouldEqual, 200)
			})

			Convey("and deleting it again should return 404 Not Found", func() {
				recorder := httptest.NewRecorder()
				request, _ := http.NewRequest("DELETE", todoTestApp.GetEndpointWithPrefix("todos")+"/"+testTodo.Id, nil)

				todoTestApp.RootRouter.ServeHTTP(recorder, request)

				So(recorder.Code, ShouldEqual, 404)
			})

			Convey("and accessing the object should gives 404 Not Found", func() {
				recorder := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", todoTestApp.GetEndpointWithPrefix("todos")+"/"+testTodo.Id, nil)

				todoTestApp.RootRouter.ServeHTTP(recorder, request)

				So(recorder.Code, ShouldEqual, 404)
			})

		})
	})
}
