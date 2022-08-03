package suggestion

import (
	"go-clean-architecture/internal/presenters/suggestion"
	"net/http"
)

func (h *handler) GetPopularKeyword(w http.ResponseWriter, r *http.Request) {
	presenter := suggestion.NewPopularKeywordPresenter(w)

	// POCO => POJO

	// factory
	popularKeywordUsecase := h.interactors.popularKeyword.GetUsecase(presenter)

	popularKeywordUsecase.Get(r.FormValue(`keyword`))

	// tell don't ask
}
