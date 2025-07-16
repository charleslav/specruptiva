package service

import (

	"cuelang.org/go/cue/cuecontext"
  "cuelang.org/go/cue/load"
	"cuelang.org/go/encoding/yaml"
)

type ValidateService struct {}

func NewValidateService() *ValidateService {
	 return &ValidateService{}
}

func (validator ValidateService) ValidateFromFiles(schema string, data string)(error) {

  ctx := cuecontext.New()
  insts := load.Instances([]string{schema}, nil)
	v := ctx.BuildInstance(insts[0])

	ymlFile, _ := yaml.Extract(data, nil)

	merged := v.Unify(ctx.BuildFile(ymlFile))

  err:= merged.Validate();

return err


}

