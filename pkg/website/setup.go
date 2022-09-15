package website

import (
	"encoding/json"

	"github.com/kyberbits/kyberbits.com/forge"
)

// Setup is foobar
func Setup(runtime *forge.Runtime) (forge.App, error) {
	app := &App{
		runtime: runtime,
		logger: &forge.LoggerJSON{
			Encoder: json.NewEncoder(runtime.Stderr),
		},
	}

	config, err := buildConfig(runtime.Environment)
	if err != nil {
		return nil, err
	}

	app.config = config

	return app, nil
}
