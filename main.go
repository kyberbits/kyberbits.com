package main

import (
	"github.com/kyberbits/kyberbits.com/forge"
	"github.com/kyberbits/kyberbits.com/pkg/website"
)

func main() {
	runtime := forge.NewRuntime()

	if err := forge.EnvironmentReadFromDefaultFiles(runtime.Environment); err != nil {
		panic(err)
	}

	app, err := website.Setup(runtime)
	if err != nil {
		panic(err)
	}

	forge.Run(app)
}
