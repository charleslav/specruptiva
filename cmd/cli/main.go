package main
// todo: utiliser adéquatement la librairie cobra et peut-être même viper
import (
	"os"
	"fmt"
	"io"
	"encoding/json"
	"github.com/spf13/cobra"
  "disruptiva.org/specruptiva/pkg/core/service"
	"disruptiva.org/specruptiva/adapters/cue"
  "disruptiva.org/specruptiva/adapters/sqlite"

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

var (
	SPECRUPTIVA_DB_FILE string
	sqliteConfig sqlite.SqliteConfig
)

func init() {

  SPECRUPTIVA_DB_FILE = "./data.db"

	sqliteConfig = sqlite.SqliteConfig{
		DbFile: SPECRUPTIVA_DB_FILE,
		LogMode: false,
	}
  var schemaStore = sqlite.NewSchemaStore(sqliteConfig)
	var schemaService = service.NewSchemaService(schemaStore)

  validateCmd := &cobra.Command{
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


schemaCmd := &cobra.Command{
    Use: "schema",
    Short: "Gère les schemas (cue)",
}

schemaCreateCmd := &cobra.Command{
    Use: "create",
    Short: "Crée un nouveau schema",
    Run: func(cmd *cobra.Command, args []string) {
	    var (
		    schema string
	    )
	    if len(args) == 0 { 
	      stdin, err := io.ReadAll(os.Stdin)
  	    if err != nil {
	  	    fmt.Println(err)
					os.Exit(1)
  	    }
		    schema=string(stdin)		 
      }else if len(args) == 1 {
		    buf, err:=os.ReadFile(args[0])
        if err != nil {
          fmt.Println(err)
	  	    os.Exit(1)
        }
	    	schema=string(buf)
      	} else {
		    fmt.Println("Erreur: il y a trop d'arguments\n   spectruptiva schemat create [SCHEMA_FILE]")
		    os.Exit(1)
    	}
    	success, err:= schemaService.Create(schema)
	    if err !=nil {
	     fmt.Println(err)
	     os.Exit(1)
  	  }
  	  output, err := json.MarshalIndent(success, "", "  ")
      if err != nil {
        fmt.Println(err)
		    os.Exit(1)
      }
      fmt.Print(string(output))
    },
}

schemaUpdateCmd := &cobra.Command{
    Use: "update",
    Short: "Modifier un schema existant",
    Run: func(cmd *cobra.Command, args []string) {
	    var (
		    schema string
	    )
	    if len(args) == 1 { 
	      stdin, err := io.ReadAll(os.Stdin)
  	    if err != nil {
	  	    fmt.Println(err)
					os.Exit(1)
  	    }
		    schema=string(stdin)		 
      }else if len(args) == 2 {
		    buf, err:=os.ReadFile(args[1])
        if err != nil {
          fmt.Println(err)
	  	    os.Exit(1)
        }
	    	schema=string(buf)
      	} else {
		    fmt.Println("Erreur: il y a trop d'arguments\n   spectruptiva schema update SCHEMA_ID [SCHEMA_FILE]")
		    os.Exit(1)
    	}
    	success, err:= schemaService.Update(args[0],schema)
	    if err !=nil {
	     fmt.Println(err)
	     os.Exit(1)
  	  }
  	  output, err := json.MarshalIndent(success, "", "  ")
      if err != nil {
        fmt.Println(err)
		    os.Exit(1)
      }
      fmt.Print(string(output))
    },
}
schemaListCmd := &cobra.Command{
    Use: "list",
    Short: "Afficher la liste des schemas existants",
    Run: func(cmd *cobra.Command, args []string) {
	    if len(args) > 0 { 
		    fmt.Println("Erreur: il y a trop d'arguments\n   spectruptiva schema list")
		    os.Exit(1)
      }
    	result, err:= schemaService.List()
	    if err !=nil {
	     fmt.Println(err)
	     os.Exit(1)
  	  }
  	  output, err := json.MarshalIndent(result, "", "  ")
      if err != nil {
        fmt.Println(err)
		    os.Exit(1)
      }
      fmt.Print(string(output))
    },
}

schemaDeleteCmd := &cobra.Command{
    Use: "delete",
    Short: "Supprime un schema existant",
    Run: func(cmd *cobra.Command, args []string) {
	    if len(args) != 1 { 
		    fmt.Println("Erreur: mauvais nombre d'arguments\n   spectruptiva schema delete SCHEMA_ID")
		    os.Exit(1)
      }
    	success, err:= schemaService.Delete(args[0])
	    if err !=nil {
	     fmt.Println(err)
	     os.Exit(1)
  	  }
  	  output, err := json.MarshalIndent(success, "", "  ")
      if err != nil {
        fmt.Println(err)
		    os.Exit(1)
      }
      fmt.Print(string(output))
    },
}

schemaCmd.AddCommand(schemaCreateCmd,schemaUpdateCmd,schemaListCmd, schemaDeleteCmd)
rootCmd.AddCommand(validateCmd, schemaCmd)

}


func main() {
	Execute()
}

