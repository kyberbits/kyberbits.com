package website

import (
	"fmt"
	"net/http"

	"github.com/kyberbits/kyberbits.com/forge"
	"github.com/kyberbits/kyberbits.com/resources"
)

// App is foobar
type App struct {
	runtime *forge.Runtime
	logger  forge.Logger
	config  *Config
}

// ListenAddress is foobar
func (app *App) ListenAddress() string {
	return fmt.Sprintf("%s:%d", app.config.Host, app.config.Port)
}

// Logger is foobar
func (app *App) Logger() forge.Logger {
	return app.logger
}

// Background is foobar
func (app *App) Background() {
}

// Handler is foobar
func (app *App) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/api/", &forge.Router{
		NotFoundHandler: app.apiHandlerNotFound(),
		Routes: map[string]http.Handler{
			"/api/greeting": app.apiHandlerGreeting(),
		},
	})

	mux.Handle("/", &forge.Static{
		FileSystem: http.FS(resources.Public),
	})

	return mux
}
