package main

import (
	"os"
	"io"
	"fmt"
  "disruptiva.org/specruptiva/pkg/core/service"
	"disruptiva.org/specruptiva/adapters/cue"
)

  var (
		validator service.ValidateService
		schema string
		data string
	)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("ValidateServiceur: il manque des arguments\n   spectruptiva SCHEMA_FILE [DATA_FILE]")
		os.Exit(1)
	}

	if len(os.Args) == 2 { 
	  stdin, err := io.ReadAll(os.Stdin)
  	if err != nil {
	  	panic(err)
  	}
		data=string(stdin)
    // todo: check it is yaml
		// todo: traiter bytes au lieu de retransformer en string
		buf, _:=os.ReadFile(os.Args[1])
		//todo: check err...
		schema=string(buf)
		validator:= cue.NewValidator(cue.ValidateFromString)

    vs:=service.NewValidateService(validator)

	  vs.SetSchema(schema)
	  vs.SetData(data)
  	err= vs.Validate()
	  if err !=nil {
	  	fmt.Println(err)
	  	os.Exit(1)
  	}

  }else {
		schema=os.Args[1]
		data=os.Args[2]
		validator:= cue.NewValidator(cue.ValidateFromFile)
    vs:=service.NewValidateService(validator)

  	vs.SetSchema(schema)
  	vs.SetData(data)
  	err:= vs.Validate()
  	if err !=nil {
  		fmt.Println(err)
  		os.Exit(1)
  	}
	}
}
