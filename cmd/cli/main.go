package main

import (
	"os"
	"fmt"
  "disruptiva.org/specruptiva/pkg/core/service"
	"disruptiva.org/specruptiva/adapters/cue"
)

  var (
		validator *service.ValidateService
	)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Erreur: il manque des arguments\n   spectruptiva SCHEMA_FILE DATA_FILE")
		os.Exit(1)
	}

  validator:=cue.NewValidator(cue.ValidateFromFile)
  vs:=service.NewValidateService(validator)

	vs.SetSchema(os.Args[1])
	vs.SetData(os.Args[2])
	err:= vs.Validate()
	if err !=nil {
		fmt.Println(err)
		os.Exit(1)
	}

	

}
