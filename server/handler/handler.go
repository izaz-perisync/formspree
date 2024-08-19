package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/perisynctechnologies/formSpree/service"
)

type IHandler interface {
	HandleProjectSetUp(w http.ResponseWriter, r *http.Request)
	HandleDeleteProject(w http.ResponseWriter, r *http.Request)
	HandleCreateForm(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	s service.IService
}

func New(s service.IService) IHandler {
	return &Handler{
		s: s,
	}
}

func (h *Handler) HandleProjectSetUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	b := service.Project{}
	if err := readJson(r, &b); err != nil {
		log.Println(err)
		WriteJson(w, http.StatusBadRequest, service.ErrMalformedRequest)
		return
	}

	if err := h.s.ProjectSetUp(ctx, b); err != nil {
		log.Println(err)
		WriteJson(w, http.StatusBadRequest, err)
		return
	}

	WriteJson(w, http.StatusOK, nil)

}

func (h *Handler) HandleDeleteProject(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	f := service.Filter{}
	if err := readQueryParams(r, &f); err != nil {
		log.Println(err)
		WriteJson(w, http.StatusBadRequest, service.ErrMalformedRequest)
	}

	if err := h.s.DeleteProject(ctx, f); err != nil {
		log.Println(err)
		WriteJson(w, http.StatusBadRequest, err)
		return
	}

	WriteJson(w, http.StatusOK, nil)

}

func (h *Handler) HandleCreateForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	f := service.Form{}
	if err := readJson(r, &f); err != nil {
		log.Println(err)
		WriteJson(w, http.StatusBadRequest, service.ErrMalformedRequest)
	}

	if err := h.s.CreatedForm(ctx, f); err != nil {
		log.Println(err)
		WriteJson(w, http.StatusBadRequest, err)
		return
	}

	WriteJson(w, http.StatusOK, nil)
}

func readQueryParams(r *http.Request, v any) error {
	if err := schema.NewDecoder().Decode(v, r.URL.Query()); err != nil {
		return err
	}

	return nil

}

func readJson(r *http.Request, v any) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return nil

}

func WriteJson(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
