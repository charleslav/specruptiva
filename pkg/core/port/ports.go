package ports

import (
	"disruptiva.org/specruptiva/pkg/core/domain"
)

type Validator interface {
	Validate() error
	SetSchema(schema string) error
	SetData(data string) error
}

type SchemaStore interface {
	Init() error
	List() (domain.Schemas,error)
  Create(schema string) (domain.Success, error)  // <<< pas certain
	Read(id string) (domain.Schema, error)
	Update(id string,schema string) (domain.Success, error)
	Delete(id string)(domain.Success, error)
}
