package ports

import (
	"disruptiva.org/specruptiva/pkg/core/domain"
)

type Validator interface {
	Validate() error
	SetSchema(schema string) error
	SetData(data string) error
}

type DataStore interface {
	List() (domain.Datas, error)
	Create(data string) (domain.Success, error)
	Read(id string) (domain.Data, error)
	Update(id string, data string) (domain.Success, error)
	Delete(id string) (domain.Success, error)
}

type SchemaStore interface {
	List() (domain.Schemas, error)
	Create(schema string) (domain.Success, error)
	Read(id string) (domain.Schema, error)
	Update(id string, schema string) (domain.Success, error)
	Delete(id string) (domain.Success, error)
}
