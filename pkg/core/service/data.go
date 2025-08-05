package service

import (
	"disruptiva.org/specruptiva/pkg/core/domain"
	"disruptiva.org/specruptiva/pkg/core/port"
)

type DataService struct {
	store ports.DataStore
}

func NewDataService(store ports.DataStore) *DataService {
	return &DataService{store: store}
}

func (s *DataService) List() (domain.Datas, error) {
	return s.store.List()
}

func (s *DataService) Create(data string) (domain.Success, error) {
	return s.store.Create(data)
}
func (s *DataService) Read(id string) (domain.Data, error) {
	return s.store.Read(id)
}
func (s *DataService) Update(id string, data string) (domain.Success, error) {
	return s.store.Update(id, data)
}
func (s *DataService) Delete(id string) (domain.Success, error) {
	return s.store.Delete(id)
}
