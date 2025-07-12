package main

import (
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
	"cuelang.org/go/encoding/yaml"
	"log"
)

func main() {
	ctx := cuecontext.New()

	insts := load.Instances([]string{"pets.cue"}, nil)
	schema := ctx.BuildInstance(insts[0])

	ymlFile, _ := yaml.Extract("charlie.yml", nil)

	merged := schema.Unify(ctx.BuildFile(ymlFile))

	err := merged.Validate();

	if err != nil {
		log.Fatalf("Validation failed: %v", err)
	}
}
