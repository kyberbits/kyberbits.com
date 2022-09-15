package website

import (
	"net/http"

	"github.com/kyberbits/kyberbits.com/forge"
)

// APIResponse is foobar
type APIResponse struct {
	Message string `json:"message"`
}

func (app *App) apiHandlerNotFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		forge.RespondJSON(w, http.StatusNotFound, APIResponse{
			Message: "Not Found",
		})
	})
}

func (app *App) apiHandlerGreeting() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		forge.RespondJSON(w, http.StatusOK, APIResponse{
			Message: "Hello there.",
		})
	})
}
