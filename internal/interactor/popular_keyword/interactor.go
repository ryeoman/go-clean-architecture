package popularkeyword

import "go-clean-architecture/internal/entities/suggestion"

//go:generate mockgen -source=./interactor.go -destination=./interactor_mock.go -package=popularkeyword

// Presenter is abstraction for popular keyword presenter
type Presenter interface {
	Present(popularKeyword suggestion.PopularKeyword, err error)
}

// PopularKeyword will act as popular keyword usecase interactor
type interactor struct{}

func NewInteractor() *interactor {
	return &interactor{}
}

func (p *interactor) GetUsecase(presenter Presenter) UsecaseItf {
	return newUsecase(presenter)
}
