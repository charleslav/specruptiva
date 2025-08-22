package service

import (
	"disruptiva.org/specruptiva/pkg/core/port"
	"disruptiva.org/specruptiva/pkg/core/domain"
)

type SchemaService struct {
	store ports.SchemaStore
}

func NewSchemaService (store ports.SchemaStore) *SchemaService {
	return &SchemaService { store: store }
}

//func(s *SchemaService) Init() error {
//	err:= s.store.Init()
//	if err != nil {
//		return err
//	}
//	return nil
//}

func(s *SchemaService) List() (domain.Schemas, error) {
	return s.store.List()
}

func(s *SchemaService) Create(schema string, apiVersion string, kind string) (domain.Success, error) {
  return s.store.Create(schema, apiVersion, kind)
}
func(s *SchemaService) Read(id string) (domain.Schema, error) {
  return s.store.Read(id)
}
func(s *SchemaService) Update(id string, schema string) (domain.Success, error) {
  return s.store.Update(id, schema)
}
func(s *SchemaService) Delete(id string) (domain.Success, error) {
  return s.store.Delete(id)
}
