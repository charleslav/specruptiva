package main

import (
	"os"
	"fmt"
  "disruptiva.org/specruptiva/pkg/core/service"
	"disruptiva.org/specruptiva/adapters/cli"
)

  var (
		validator *service.ValidateService
	)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Erreur: il manque des arguments\n   spectruptiva SCHEMA_FILE DATA_FILE")
		os.Exit(1)
	}

  cliValidator:=cli.NewValidator()
  validator:=service.NewValidateService(&cliValidator)

	validator.SetSchema(os.Args[1])
	validator.SetData(os.Args[2])
	err:= validator.Validate()
	if err !=nil {
		fmt.Println(err)
		os.Exit(1)
	}

	

}
