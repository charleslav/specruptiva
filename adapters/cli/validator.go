
package cli

import (

	"cuelang.org/go/cue/cuecontext"
  "cuelang.org/go/cue/load"
	"cuelang.org/go/encoding/yaml"
)

func NewValidator() CliValidator {
	return CliValidator{}
}

type CliValidator struct {
	schemafile string
	datafile string
}

func (cs *CliValidator) SetSchema(schema string) error {
	// todo: check that it is a valid cue schema
	cs.schemafile = schema
	return nil
}

func (cs *CliValidator) SetData(data string) error {
	// todo: check that it is a valid cue schema
	cs.datafile = data
	return nil
}

func (cs *CliValidator) Validate() error {

	//todo: check schema and data are available

  ctx := cuecontext.New()
  insts := load.Instances([]string{cs.schemafile}, nil)
	schema := ctx.BuildInstance(insts[0])

	ymlFile, _ := yaml.Extract(cs.datafile, nil)

	merged := schema.Unify(ctx.BuildFile(ymlFile))

  err:= merged.Validate();

  return err
 }
