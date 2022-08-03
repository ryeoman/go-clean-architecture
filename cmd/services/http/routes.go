package http

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func (h *httpService) Register() {
	h.router.Use(middleware.Logger)

	h.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	h.router.Get("/keyword/popular", h.handlers.PopularKeyword.GetPopularKeyword)
}
