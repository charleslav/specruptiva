package main

import (
	"os"
	"fmt"
	"io"
	"github.com/spf13/cobra"
  "disruptiva.org/specruptiva/pkg/core/service"
	"disruptiva.org/specruptiva/adapters/cue"

)

var rootCmd = &cobra.Command{
	Use:   "specruptiva",
	Short: "Cli permermettant d'interagir avec les données de specruptiva",
	Long: "Cli permermettant d'interagir avec les données de specruptiva",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {


versionCmd := &cobra.Command{
    Use: "validate",
    Short: "Vérifie qu'une donnée (yaml) est conforme à un schema (cue)",
    Run: func(cmd *cobra.Command, args []string) {
	    var (
		    schema string
		    data string
	    )
 	    if len(args) < 1 {
		    fmt.Println("ValidateServiceur: il manque des arguments\n   spectruptiva validate SCHEMA_FILE [DATA_FILE]")
		    os.Exit(1)
    	}
	    if len(args) == 1 { 
	      stdin, err := io.ReadAll(os.Stdin)
  	    if err != nil {
	  	    panic(err)
  	    }
		    data=string(stdin)
        // todo: check it is yaml
		    // todo: traiter bytes au lieu de retransformer en string
		    buf, _:=os.ReadFile(args[0])
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
	    	schema=args[0]
		    data=args[1]
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
    },
}

sayHelloCmd := &cobra.Command{
    Use: "sayhello",
    Short: "Say Hello",
    Run: func(cmd *cobra.Command, args []string) {
        
    },
}

rootCmd.AddCommand(versionCmd, sayHelloCmd)

//	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func main() {
	Execute()
}

