package suggestion

import (
	popularkeyword "go-clean-architecture/internal/interactor/popular_keyword"
)

type interactors struct {
	popularKeyword popularKeywordInteractor
}

//go:generate mockgen -source=./handler.go -destination=./handler_mock.go -package=suggestion

type popularKeywordInteractor interface {
	GetUsecase(p popularkeyword.Presenter) popularkeyword.UsecaseItf
}

type handler struct {
	interactors interactors
}

func NewHandler(pk popularKeywordInteractor) *handler {
	return &handler{
		interactors: interactors{
			popularKeyword: pk,
		},
	}
}
