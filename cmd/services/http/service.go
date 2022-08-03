package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type suggestionHandlers interface {
	GetPopularKeyword(w http.ResponseWriter, r *http.Request)
}

type Handlers struct {
	PopularKeyword suggestionHandlers
}

type httpService struct {
	router *chi.Mux

	handlers Handlers
}

func NewService(h Handlers) *httpService {
	return &httpService{
		router:   chi.NewRouter(),
		handlers: h,
	}
}

func (h *httpService) Start() {
	http.ListenAndServe(":3000", h.router)
}
