package main

import (
	"fmt"
	"time"
	// ONLY FOR TEST
	"log"
	// ONLY FOR TEST
	"reflect"
)

// Custom error type for when an entry is not found
type RowNotFoundError string

func (e RowNotFoundError) Error() string {
	return fmt.Sprintf("The object with ID '%s' could not be found", string(e))
}

// Interface for data objects with getters/setters for Id, Created fields
type EntityHandler interface {
	GetId() string
	SetId(id string)
	GetCreated() time.Time
	SetCreated(t time.Time)
}

// Base data object that implements the EntityHandler interface, to be embedded into concrete structs
type Entity struct {
	Id      string    `json:"id"`
	Created time.Time `json:"created"`
}

func (d *Entity) GetId() string          { return d.Id }
func (d *Entity) SetId(id string)        { d.Id = id }
func (d *Entity) GetCreated() time.Time  { return d.Created }
func (d *Entity) SetCreated(t time.Time) { d.Created = t }

// Interface for DB service structs
type DataAccessor interface {
	Create(src interface{}) (id string, err error)
	Delete(id string) (err error)
	Read(id string, dest interface{}) (err error)
	ReadMany(dests interface{}) (err error)
	Update(id string, srcToUpdate interface{}) (err error)
}

// Base DB service that implement most of the DataAccessor interface (except for GetTblName()), should be embedded into concreate DB services
type DBService struct {
	TblName string
}

// ONLY FOR TEST
var mockRepo = make(map[string]interface{})

// ONLY FOR TEST
func GetUUID() string {
	return "3644ad8c-27cd-43f8-b05a-ff272b4e4808"
}

// Insert src object to db
func (s *DBService) Create(src interface{}) (id string, err error) {

	// ONLY FOR TEST
	id = GetUUID()

	val := reflect.ValueOf(src)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	v := val.FieldByName("Id")
	v.SetString(id)

	log.Println(src)
	mockRepo[id] = src

	return
}

// Delete entry from db
func (s *DBService) Delete(id string) (err error) {

	// ONLY FOR TEST
	if _, ok := mockRepo[id]; ok {
		delete(mockRepo, id)
	} else {
		return RowNotFoundError(id)
	}

	return
}

// Read entry by id from db
func (s *DBService) Read(id string, dest interface{}) (err error) {

	// ONLY FOR TEST
	if _, ok := mockRepo[id]; ok {
		dest = mockRepo[id]
	} else {
		return RowNotFoundError(id)
	}

	return
}

// Read entries from db
func (s *DBService) ReadMany(dests interface{}) (err error) {

	// ONLY FOR TEST
	var tmp []interface{}
	for _, val := range mockRepo {
		tmp = append(tmp, val)
	}
	dests = tmp

	return
}

// Update entry in db from srcToUpdate
func (s *DBService) Update(id string, srcToUpdate interface{}) (err error) {

	// ONLY FOR TEST
	if _, ok := mockRepo[id]; ok {
		mockRepo[id] = srcToUpdate
	} else {
		return RowNotFoundError(id)
	}

	return
}
