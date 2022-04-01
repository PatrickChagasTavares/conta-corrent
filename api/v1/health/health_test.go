package health

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/patrickchagastavares/StoneTest/app"
	"github.com/patrickchagastavares/StoneTest/mocks"
	"github.com/patrickchagastavares/StoneTest/model"
	"github.com/patrickchagastavares/StoneTest/test"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	startedAt := time.Now()
	cases := map[string]struct {
		ExpectedResponseData *model.Health
		ExpectedErr          error
		PrepareMock          func(mock *mocks.MockHealthApp)
	}{
		"return sucess": {
			ExpectedResponseData: &model.Health{
				DatabaseStatus: "OK",
				ServerStatedAt: startedAt.UTC().String(),
			},
			PrepareMock: func(mock *mocks.MockHealthApp) {
				mock.
					EXPECT().
					Ping(gomock.Any()).
					Times(1).
					Return(&model.Health{
						DatabaseStatus: "OK",
						ServerStatedAt: startedAt.UTC().String(),
					}, nil)
			},
		},
		"return error: an error has occurred": {
			ExpectedErr: model.NewError(http.StatusInternalServerError, "ocorreu um erro", nil),
			PrepareMock: func(mock *mocks.MockHealthApp) {
				mock.
					EXPECT().
					Ping(gomock.Any()).
					Times(1).
					Return(
						nil,
						model.NewError(
							http.StatusInternalServerError,
							"ocorreu um erro",
							nil,
						),
					)
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)

			mockApp := mocks.NewMockHealthApp(ctrl)
			cs.PrepareMock(mockApp)

			h := handler{
				apps: &app.Container{
					Health: mockApp,
				},
			}

			req := httptest.NewRequest(http.MethodGet, "/v1/health", nil).WithContext(ctx)
			rec := httptest.NewRecorder()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			e := echo.New()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/health")

			err := h.ping(c)

			assert.Equal(t, cs.ExpectedErr, err)

			var resp model.Response
			var respData *model.Health
			json.NewDecoder(rec.Body).Decode(&resp)

			dataJson, _ := json.Marshal(resp.Data)
			json.Unmarshal(dataJson, &respData)

			assert.Equal(t, cs.ExpectedResponseData, respData)

		})
	}

}
