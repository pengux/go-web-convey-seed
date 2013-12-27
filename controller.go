package main

import (
	"encoding/json"
	"github.com/gocraft/web"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Base controller struct with common methods, to be embedded into concrete structs
type Controller struct {
}

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

// Create action, object's Id and Created field are set to default values
func (c *Controller) Create(da DataAccessor, rw web.ResponseWriter, object EntityHandler, endpoint string) {

	object.SetId("")
	object.SetCreated(time.Now())

	id, err := da.Create(object)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Panicln(err)
	}

	rw.Header().Set("Location", endpoint+"/"+id)
	rw.WriteHeader(http.StatusCreated)
}

// Delete action
func (c *Controller) Delete(da DataAccessor, rw web.ResponseWriter, id string) {

	var object interface{}
	err := da.Read(id, &object)
	if err != nil {
		if _, ok := err.(RowNotFoundError); ok {
			rw.WriteHeader(http.StatusNotFound)
			return
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
			log.Panicln(err)
		}
	}

	err = da.Delete(id)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		log.Panicln(err)
	}

	rw.WriteHeader(http.StatusOK)
}

// Read action by id, the JSON response that is returned is indented with json.MarshalIndent
func (c *Controller) Read(da DataAccessor, rw web.ResponseWriter, id string, object interface{}) {

	err := da.Read(id, object)
	if err != nil {
		if _, ok := err.(RowNotFoundError); ok {
			rw.WriteHeader(http.StatusNotFound)
			return
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
			log.Panicln(err)
		}
	}

	resp, err := json.MarshalIndent(object, "", "    ")
	if err != nil {
		log.Panicln(err)
	}

	rw.Write(resp)
}

// Read multiple entries, the JSON response that is returned is indented with json.MarshalIndent
func (c *Controller) ReadMany(da DataAccessor, rw web.ResponseWriter, objects interface{}) {

	err := da.ReadMany(objects)
	log.Println(objects)
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

// Replace action, replace the whole object
func (c *Controller) Replace(da DataAccessor, rw web.ResponseWriter, id string, object EntityHandler) {

	object.SetId("")
	object.SetCreated(time.Time{})

	err := da.Update(id, object)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		log.Panicln(err)
	}

	rw.WriteHeader(http.StatusOK)
}
