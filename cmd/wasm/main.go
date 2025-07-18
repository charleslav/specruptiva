package main

import (
	"syscall/js"
	"fmt"
	"disruptiva.org/specruptiva/pkg/core/service"
	"disruptiva.org/specruptiva/adapters/cue"
)


var (
	validator service.ValidateService
	schema string
	data string
)



func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
        
    if len(args) !=2 {
			return "2 arguments sont requis. (SCHEMA , DATA)"
		}
    // todo: mettre tout ça dans une fonciton à part
		validator:= cue.NewValidator(cue.ValidateFromString)
    vs:=service.NewValidateService(validator)

    schema=args[0].String()
    data=args[1].String()
		fmt.Println("schema: ", schema)
		fmt.Println("data: ", data)

  	vs.SetSchema(schema)
  	vs.SetData(data)
  	err:= vs.Validate()
  	if err !=nil {
  		return err.Error()
  	}
		return 0

	})
	return jsonFunc
}

func main(){
	fmt.Println("hello from go wasm")
  js.Global().Set("validate", jsonWrapper())
	<-make(chan struct{})
}
