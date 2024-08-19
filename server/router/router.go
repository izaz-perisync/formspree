package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/perisynctechnologies/formSpree/server/handler"
)

func BuildRoute(h handler.IHandler) *mux.Router {

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1/fs").Subrouter()
	project := api.PathPrefix("/project").Subrouter()
	project.HandleFunc("/manage", h.HandleProjectSetUp).Methods(http.MethodPost)
	project.HandleFunc("/delete", h.HandleDeleteProject).Methods(http.MethodDelete)
	form := project.PathPrefix("/form").Subrouter()
	form.HandleFunc("/create", h.HandleCreateForm).Methods(http.MethodPost)
	return r
}
