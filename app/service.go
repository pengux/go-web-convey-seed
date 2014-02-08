package app

import (
	"fmt"
	"time"
)

// Custom error type for when an entry is not found
type RowNotFoundError string

func (e RowNotFoundError) Error() string {
	return fmt.Sprintf("The object with ID '%s' could not be found", string(e))
}

// Interface for data objects with getters/setters for Id, CreatedAt fields
type EntityHandler interface {
	GetId() string
	SetId(id string)
	GetCreatedAt() time.Time
	SetCreatedAt(t time.Time)
}

type EntitiesFactory interface {
	Append(object interface{})
}

type DataAccessor interface {
	Create(object EntityHandler) (id string, err error)
	Delete(id string) error
	Read(id string, object EntityHandler) (result EntityHandler, err error)
	ReadMany(objects EntitiesFactory) (err error)
	Update(id string, object EntityHandler) error
}
