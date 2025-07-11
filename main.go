package main

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/errors"
	"cuelang.org/go/cue/load"
	"fmt"
	"log"
)

func main() {
	ctx := cuecontext.New()

	// Load the package "example" from the current directory.
	insts := load.Instances([]string{"."}, nil)

	// Build the instance
	v := ctx.BuildInstance(insts[0])

	// Check for errors and print detailed information
	if err := v.Err(); err != nil {
		fmt.Println("=== CUE VALIDATION ERRORS ===")

		// Use errors.Details for the cleanest, most comprehensive output
		fmt.Println(errors.Details(err, nil))

		log.Fatal("CUE validation failed")
	}

	// If no errors, proceed with lookup
	output := v.LookupPath(cue.ParsePath("output"))
	if output.Err() != nil {
		fmt.Println("Error looking up 'output' field:")
		fmt.Println(errors.Details(output.Err(), nil))
		return
	}

	fmt.Println("Output:", output)
}
