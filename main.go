package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jingweno/conf"
	"github.com/pengux/web"
)

type Context struct{}

// Unmarshal request body in JSON to a struct
func (c *Context) UnmarshalBody(rw web.ResponseWriter, req *web.Request, object interface{}) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Panicln(err)
	}

	err = json.Unmarshal(body, object)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Panicln(err)
	}
}

func main() {
	defaults := map[string]interface{}{
		"ENV":        "production",
		"API_PORT":   9090,
		"API_PREFIX": "",
	}
	config, err := conf.NewLoader().Env().Argv().File("./config.json").Defaults(defaults).Load()
	if err != nil {
		log.Fatalln(err)
	}

	rootRouter := web.NewWithPrefix(Context{}, config.Get("API_PREFIX").(string))

	rootRouter.Subrouter(TodoContext{}, "/todos").
		Get("/", (*TodoContext).ReadMany).
		Get("/:id", (*TodoContext).Read).
		Post("/", (*TodoContext).Create).
		Put("/:id", (*TodoContext).Replace).
		Delete("/:id", (*TodoContext).Delete).
		Middleware((*TodoContext).Init)

	log.Printf("Listen to port %v", config.Get("API_PORT"))
	http.ListenAndServe(fmt.Sprintf(":%v", config.Get("API_PORT")), rootRouter)
}
