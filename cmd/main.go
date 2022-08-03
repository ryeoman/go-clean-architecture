package main

import (
	"go-clean-architecture/cmd/services/http"
	"go-clean-architecture/internal/handlers/suggestion"
	popularkeyword "go-clean-architecture/internal/interactor/popular_keyword"
)

func main() {
	popularKeywordInteractor := popularkeyword.NewInteractor()

	handlers := http.Handlers{
		PopularKeyword: suggestion.NewHandler(popularKeywordInteractor),
	}
	httpService := http.NewService(handlers)
	httpService.Register()
	httpService.Start()
}
