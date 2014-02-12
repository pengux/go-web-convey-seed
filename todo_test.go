package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jingweno/conf"
	"github.com/pengux/web"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTodo(t *testing.T) {
	defaults := map[string]interface{}{
		"ENV":        "testing",
		"API_PORT":   9090,
		"API_PREFIX": "",
	}
	config, err := conf.NewLoader().Env().Argv().File("./config_test.json").Defaults(defaults).Load()
	if err != nil {
		log.Fatalln(err)
	}

	// Init DB service
	dbService := &InMemoryDBService{}
	dbService.Init()

	rootRouter := web.NewWithPrefix(Context{}, config.Get("API_PREFIX").(string))

	rootRouter.Subrouter(TodoContext{}, "/todos").
		Get("/", (*TodoContext).ReadMany).
		Get("/:id", (*TodoContext).Read).
		Post("/", (*TodoContext).Create).
		Put("/:id", (*TodoContext).Replace).
		Delete("/:id", (*TodoContext).Delete).
		Middleware((*TodoContext).Init)

	Convey("Given the todos endpoint", t, func() {
		var testTodo Todo

		Convey("When adding a todo to the collection", func() {
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("POST", config.Get("API_PREFIX").(string)+"/todos", strings.NewReader(`{
				"title": "Master Go",
				"done": false
			}`))

			rootRouter.ServeHTTP(recorder, request)

			Convey("It should returns a 201 Created status", func() {
				So(recorder.Code, ShouldEqual, 201)
			})

		})

		Convey("When listing all todos from the collection", func() {
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", config.Get("API_PREFIX").(string)+"/todos", nil)

			rootRouter.ServeHTTP(recorder, request)

			Convey("It should contains the 'Testtodo' object(s)", func() {
				var todos []Todo
				json.Unmarshal(recorder.Body.Bytes(), &todos)
				testTodo = todos[0]

				So(testTodo.Title, ShouldEqual, "Master Go")
			})
		})

		Convey("When updating the 'Testtodo'", func() {
			testTodo.Done = true
			data, _ := json.Marshal(testTodo)

			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("PUT", config.Get("API_PREFIX").(string)+"/todos"+"/"+testTodo.Id, strings.NewReader(string(data)))

			rootRouter.ServeHTTP(recorder, request)

			Convey("the data should persist", func() {
				recorder := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", config.Get("API_PREFIX").(string)+"/todos"+"/"+testTodo.Id, nil)

				rootRouter.ServeHTTP(recorder, request)
				var testTodoUpdated Todo
				json.Unmarshal(recorder.Body.Bytes(), &testTodoUpdated)

				So(testTodoUpdated.Done, ShouldEqual, true)
			})
		})

		Convey("When deleting the 'Testtodo'", func() {
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("DELETE", config.Get("API_PREFIX").(string)+"/todos"+"/"+testTodo.Id, nil)

			rootRouter.ServeHTTP(recorder, request)

			Convey("it should returns 200 Status", func() {
				So(recorder.Code, ShouldEqual, 200)
			})

			Convey("and deleting it again should return 404 Not Found", func() {
				recorder := httptest.NewRecorder()
				request, _ := http.NewRequest("DELETE", config.Get("API_PREFIX").(string)+"/todos/"+testTodo.Id, nil)

				rootRouter.ServeHTTP(recorder, request)

				So(recorder.Code, ShouldEqual, 404)
			})

			Convey("and accessing the object should gives 404 Not Found", func() {
				recorder := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", config.Get("API_PREFIX").(string)+"/todos/"+testTodo.Id, nil)

				rootRouter.ServeHTTP(recorder, request)

				So(recorder.Code, ShouldEqual, 404)
			})

		})
	})
}
