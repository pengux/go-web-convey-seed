package main

import "code.google.com/p/go-uuid/uuid"

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

func (s *InMemoryDBService) Read(id string, object EntityHandler) (result EntityHandler, err error) {
	if val, ok := db[s.TblName][id]; ok {
		object = val
		return object, nil
	}

	return nil, RowNotFoundError(id)
}

func (s *InMemoryDBService) ReadMany(objects EntitiesFactory) (err error) {
	for _, object := range db[s.TblName] {
		objects.Append(object)
	}
	return nil
}

func (s *InMemoryDBService) Update(id string, object EntityHandler) error {
	db[s.TblName][id] = object

	return nil
}
