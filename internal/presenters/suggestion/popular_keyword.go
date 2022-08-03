package suggestion

import (
	"encoding/json"
	"go-clean-architecture/internal/entities/suggestion"
	"net/http"
)

type popularKeywordPresenter struct {
	writer http.ResponseWriter
}

func NewPopularKeywordPresenter(w http.ResponseWriter) *popularKeywordPresenter {
	return &popularKeywordPresenter{
		writer: w,
	}
}

func (p *popularKeywordPresenter) Present(
	popularKeyword suggestion.PopularKeyword,
	err error,
) {
	if err != nil {
		p.writer.WriteHeader(http.StatusInternalServerError)
		body, _ := json.Marshal(map[string]interface{}{
			"error": err,
		})
		p.writer.Write(body)
		return
	}
	p.writer.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(map[string]interface{}{
		"keyword":  popularKeyword.Keyword,
		"priority": popularKeyword.Priority,
	})
	p.writer.Write(body)
}
