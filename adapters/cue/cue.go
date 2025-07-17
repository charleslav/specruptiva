package cue

import (
	"disruptiva.org/specruptiva/pkg/core/port"
	"cuelang.org/go/cue/cuecontext"
  "cuelang.org/go/cue/load"
	"cuelang.org/go/encoding/yaml"
)

type validateFunc func(schema string, data string) error

func NewValidator( validate validateFunc ) ports.Validator{
	return &CueValidator{ validate: validate }
}

type CueValidator struct {
	schema string
	data string
	validate validateFunc
}

func(cv *CueValidator) SetSchema(schema string) error {
	cv.schema = schema
	return nil
}

func(cv *CueValidator) SetData(data string) error {
	cv.data = data
	return nil
}

func(cv *CueValidator) Validate() error {
	
		return cv.validate(cv.schema,cv.data)
		 
}

func ValidateFromString(schema string, data string) error {

  ctx := cuecontext.New()
	v := ctx.CompileString(schema)
	merged := v.Unify(ctx.CompileString(data))
  err:= merged.Validate();
 
	return err
}

func ValidateFromFile(schemafile string, datafile string) error {
  ctx := cuecontext.New()
  insts := load.Instances([]string{schemafile}, nil)
	schema := ctx.BuildInstance(insts[0])

	ymlFile, err := yaml.Extract(datafile, nil)

	if err != nil {
		return err
	}else{

	  merged := schema.Unify(ctx.BuildFile(ymlFile))
    err:= merged.Validate();
	  return err
	}
}

