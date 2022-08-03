package suggestion

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_handler_GetPopularKeyword(t *testing.T) {
	type fields struct {
		interactor *MockpopularKeywordInteractor
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		setup  func(mock fields)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := fields{
				interactor: NewMockpopularKeywordInteractor(ctrl),
			}
			tt.setup(mock)

			h := &handler{
				interactors: interactors{
					popularKeyword: mock.interactor,
				},
			}
			h.GetPopularKeyword(tt.args.w, tt.args.r)
		})
	}
}
