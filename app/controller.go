package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gocraft/web"
)

// Base controller struct with common methods, to be embedded into concrete structs
type Controller struct{}

// Unmarshal request body in JSON to a struct
func (c *Controller) UnmarshalBody(rw web.ResponseWriter, req *web.Request, object interface{}) {
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

// Insert object into db
func (c *Controller) Create(da DataAccessor, rw web.ResponseWriter, object EntityHandler, redirectUrl string) {
	object.SetId("")
	object.SetCreatedAt(time.Now())

	id, err := da.Create(object)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Panicln(err)
	}

	rw.Header().Set("Location", redirectUrl+"/"+id)
	rw.WriteHeader(http.StatusCreated)
}

func (c *Controller) Delete(da DataAccessor, rw web.ResponseWriter, id string) {
	err := da.Delete(id)
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

func (c *Controller) Read(da DataAccessor, rw web.ResponseWriter, req *web.Request, object EntityHandler) {
	result, err := da.Read(req.PathParams["id"], object)
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

func (c *Controller) ReadMany(da DataAccessor, rw web.ResponseWriter, req *web.Request, objects EntitiesFactory) {
	err := da.ReadMany(objects)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Panicln(err)
	}

	resp, err := json.MarshalIndent(objects, "", "    ")
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Panicln(err)
	}

	rw.Write(resp)
}

func (c *Controller) Replace(da DataAccessor, rw web.ResponseWriter, id string, object EntityHandler) {
	err := da.Update(id, object)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Panicln(err)
	}

	rw.WriteHeader(http.StatusOK)
}
