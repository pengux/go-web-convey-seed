package app

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
)

// Base data object that implements the EntityHandler interface, to be embedded into concrete structs
type Entity struct {
	Id        string
	CreatedAt time.Time
}

func (d *Entity) GetId() string            { return d.Id }
func (d *Entity) SetId(id string)          { d.Id = id }
func (d *Entity) GetCreatedAt() time.Time  { return d.CreatedAt }
func (d *Entity) SetCreatedAt(t time.Time) { d.CreatedAt = t }

var (
	db = make(map[string]map[string]EntityHandler)
)

type InMemoryDBService struct {
	TblName string
}

func (s *InMemoryDBService) Init() {
	// Init table if it doesn't exists
	if _, ok := db[s.TblName]; !ok {
		db[s.TblName] = make(map[string]EntityHandler)
	}
}

func (s *InMemoryDBService) Create(object EntityHandler) (id string, err error) {
	// Set arbitary Id string
	id = uuid.New()
	object.SetId(id)
	db[s.TblName][id] = object

	return id, nil
}

func (s *InMemoryDBService) Delete(id string) error {
	if _, ok := db[s.TblName][id]; ok {
		delete(db[s.TblName], id)
	} else {
		return RowNotFoundError(id)
	}

	return nil
}

func (s *InMemoryDBService) Read(id string) (object EntityHandler, err error) {
	if val, ok := db[s.TblName][id]; ok {
		return val, nil
	}

	return nil, RowNotFoundError(id)
}

func (s *InMemoryDBService) ReadMany() (objects []EntityHandler, err error) {
	for _, object := range db[s.TblName] {
		objects = append(objects, object)
	}
	return objects, nil
}

func (s *InMemoryDBService) Update(id string, object EntityHandler) error {
	db[s.TblName][id] = object

	return nil
}
