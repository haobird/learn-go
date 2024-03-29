package main

import (
	"fmt"
	"log"

	"github.com/google/cel-go/cel"
)

func main() {
	env, err := cel.NewEnv(
		cel.Variable("name", cel.StringType),
		cel.Variable("group", cel.StringType),
	)
	ast, issues := env.Compile(`name.startsWith("/groups/" + group)`)
	if issues != nil && issues.Err() != nil {
		log.Fatalf("type-check error: %s", issues.Err())
	}

	prg, err := env.Program(ast)
	if err != nil {
		log.Fatalf("program construction error: %s", err)
	}

	out, details, err := prg.Eval(map[string]interface{}{
		"name":  "/groups/acme.co/documents/secret-stuff",
		"group": "acme.co",
	})
	fmt.Println(out)
	fmt.Println(details)
}
