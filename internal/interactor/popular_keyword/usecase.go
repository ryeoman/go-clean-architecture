package popularkeyword

import "go-clean-architecture/internal/entities/suggestion"

//go:generate mockgen -source=./usecase.go -destination=./usecase_mock.go -package=popularkeyword

type UsecaseItf interface {
	Get(keyword string)
}

type usecase struct {
	presenter Presenter
}

func newUsecase(p Presenter) *usecase {
	return &usecase{
		presenter: p,
	}
}

func (u *usecase) Get(keyword string) {
	u.presenter.Present(suggestion.PopularKeyword{
		Keyword:  "a keyword",
		Priority: 1,
	}, nil)
}
