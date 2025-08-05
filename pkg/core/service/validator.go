package service

import (
	"disruptiva.org/specruptiva/pkg/core/port"
)

type ValidateService struct {
  validator ports.Validator
}

func NewValidateService(validator ports.Validator) *ValidateService {
	return &ValidateService{ validator: validator }
}

func (vs *ValidateService) SetSchema(schema string) error {
	return vs.validator.SetSchema(schema)
} 

func (vs *ValidateService) SetData(data string) error {
	return vs.validator.SetData(data)
} 

func (vs *ValidateService) Validate() error {
	return vs.validator.Validate()
} 
