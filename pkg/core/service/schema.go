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

func(s *SchemaService) Init() error {
	err:= s.store.Init()
	if err != nil {
		return err
	}
	return nil
}

func(s *SchemaService) List() (domain.Schemas, error) {
	return nil, nil
}


