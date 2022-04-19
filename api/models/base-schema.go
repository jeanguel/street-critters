package models

import (
	"io"
)

type BaseModel interface {
	Save() error
}

type BaseSchema interface {
	Map(io.ReadCloser) (BaseSchema, error)
	Validate() []error
	ToModel() (BaseModel, error)
}
