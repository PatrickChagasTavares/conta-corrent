package login

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/patrickchagastavares/conta-corrent/app"
	"github.com/patrickchagastavares/conta-corrent/mocks"
	"github.com/patrickchagastavares/conta-corrent/model"
	"github.com/patrickchagastavares/conta-corrent/test"
	"github.com/patrickchagastavares/conta-corrent/utils/session"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {

	input := &auth{
		CPF:      "74596687021",
		Password: "123456",
	}

	cases := map[string]struct {
		ExpectedResponseData *session.SessionAuth
		ExpectedErr          error

		InputBody   func() io.Reader
		PrepareMock func(mock *mocks.MockLoginApp)
	}{
		"return sucess": {
			InputBody: func() io.Reader {
				bt, _ := json.Marshal(input)
				return bytes.NewReader(bt)
			},
			ExpectedResponseData: &session.SessionAuth{
				Token: "token",
			},
			PrepareMock: func(mock *mocks.MockLoginApp) {
				mock.EXPECT().Login(gomock.Any(), gomock.Any(), gomock.Any()).Return(&session.SessionAuth{
					Token: "token",
				}, nil)
			},
		},
		"return error: bind": {
			InputBody: func() io.Reader {
				return strings.NewReader("invalid")
			},
			ExpectedErr: errLoginBind,
			PrepareMock: func(mock *mocks.MockLoginApp) {},
		},
		"return error: application": {
			InputBody: func() io.Reader {
				bt, _ := json.Marshal(input)
				return bytes.NewReader(bt)
			},
			ExpectedErr: model.NewError(http.StatusInternalServerError, "tivemos um erro ao realizar o login", nil),
			PrepareMock: func(mock *mocks.MockLoginApp) {
				mock.EXPECT().Login(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, model.NewError(http.StatusInternalServerError, "tivemos um erro ao realizar o login", nil))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)

			mockApp := mocks.NewMockLoginApp(ctrl)
			cs.PrepareMock(mockApp)
			h := handler{
				apps: &app.Container{
					Login: mockApp,
				},
			}

			req := httptest.NewRequest(http.MethodPost, "/v1/login", cs.InputBody()).WithContext(ctx)
			rec := httptest.NewRecorder()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			e := echo.New()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/login")

			err := h.login(c)

			assert.Equal(t, cs.ExpectedErr, err)

			var resp model.Response
			json.NewDecoder(rec.Body).Decode(&resp)

			var respData *session.SessionAuth
			respDataJson, _ := json.Marshal(resp.Data)
			json.Unmarshal(respDataJson, &respData)

			assert.Equal(t, cs.ExpectedResponseData, respData)

		})
	}

}
