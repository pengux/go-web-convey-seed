package main

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFoo(t *testing.T) {

	app := Bootstrap()

	Convey("Given the foos endpoint", t, func() {
		var testFoo Foo

		Convey("When adding an foo to the collection", func() {
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("POST", Endpoints["foos"], strings.NewReader(`{
	    "name": "Testfoo",
	    "phone": "123"
	}`))

			app.RootRouter.ServeHTTP(recorder, request)

			Convey("It should returns a 201 Created status", func() {
				So(recorder.Code, ShouldEqual, 201)
			})

		})

		Convey("When listing all foos from the collection", func() {
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", Endpoints["foos"], nil)

			app.RootRouter.ServeHTTP(recorder, request)

			Convey("It should contains the 'Testfoo' object(s)", func() {
				var foos []Foo
				log.Println(string(recorder.Body.Bytes()))
				err := json.Unmarshal(recorder.Body.Bytes(), &foos)
				if err != nil {
					panic(err)
				}

				for _, foo := range foos {
					if foo.Name == "Testfoo" {
						testFoo = foo
						break
					}
				}

				So(testFoo.Id, ShouldNotEqual, "")
			})
		})

		Convey("When updating the 'Testfoo'", func() {
			testFoo.Phone = "abc"
			data, err := json.Marshal(testFoo)
			if err != nil {
				panic(err)
			}

			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("PUT", Endpoints["foos"]+"/"+testFoo.Id, strings.NewReader(string(data)))

			app.RootRouter.ServeHTTP(recorder, request)

			Convey("the data should persist", func() {
				recorder := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", Endpoints["foos"]+"/"+testFoo.Id, nil)

				app.RootRouter.ServeHTTP(recorder, request)
				var testFooUpdated Foo
				err := json.Unmarshal(recorder.Body.Bytes(), &testFooUpdated)
				if err != nil {
					panic(err)
				}

				So(testFooUpdated.Phone, ShouldEqual, "abc")
			})
		})

		Convey("When deleting the 'Testfoo'", func() {
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("DELETE", Endpoints["foos"]+"/"+testFoo.Id, nil)

			app.RootRouter.ServeHTTP(recorder, request)

			Convey("it should returns 200 Status", func() {
				So(recorder.Code, ShouldEqual, 200)
			})

			Convey("and deleting it again should return 404 Not Found", func() {
				recorder := httptest.NewRecorder()
				request, _ := http.NewRequest("DELETE", Endpoints["foos"]+"/"+testFoo.Id, nil)

				app.RootRouter.ServeHTTP(recorder, request)

				So(recorder.Code, ShouldEqual, 404)
			})

			Convey("and accessing the object should gives 404 Not Found", func() {
				recorder := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", Endpoints["foos"]+"/"+testFoo.Id, nil)

				app.RootRouter.ServeHTTP(recorder, request)

				So(recorder.Code, ShouldEqual, 404)
			})

		})
	})
}
