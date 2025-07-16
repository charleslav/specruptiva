package main

import (
	"os"
	"fmt"
  "disruptiva.org/specruptiva/pkg/core/service"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Erreur: il manque des arguments\n   spectruptiva SCHEMA_FILE DATA_FILE")
		os.Exit(1)
	}
  
  validator:=service.NewValidateService()

	err:= validator.ValidateFromFiles(os.Args[1], os.Args[2])
	if err !=nil {
		fmt.Println(err)
		os.Exit(1)
	}

	

}
