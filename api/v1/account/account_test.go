package account

import (
	"bytes"
	"encoding/json"
	"io"
	"math/big"
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
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	lists := []*model.Account{
		{ID: 1, Name: "test", CPF: "", Balance: *big.NewInt(100)},
	}

	cases := map[string]struct {
		ExpectedResponseData []*model.Account
		ExpectedErr          error

		PrepareMock func(mock *mocks.MockAccountApp)
	}{
		"return sucess": {
			ExpectedResponseData: lists,
			PrepareMock: func(mock *mocks.MockAccountApp) {
				mock.EXPECT().List(gomock.Any()).Return(lists, nil)
			},
		},
		"return error: application ": {
			ExpectedErr: model.NewError(http.StatusInternalServerError, "error", nil),
			PrepareMock: func(mock *mocks.MockAccountApp) {
				mock.EXPECT().List(gomock.Any()).Return(nil, model.NewError(http.StatusInternalServerError, "error", nil))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)

			mockApp := mocks.NewMockAccountApp(ctrl)
			cs.PrepareMock(mockApp)

			h := handler{
				apps: &app.Container{
					Account: mockApp,
				},
			}

			req := httptest.NewRequest(http.MethodGet, "/v1/account", nil).WithContext(ctx)
			rec := httptest.NewRecorder()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			e := echo.New()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/account")

			err := h.list(c)

			assert.Equal(t, cs.ExpectedErr, err)

			var resp model.Response
			json.NewDecoder(rec.Body).Decode(&resp)

			var respData []*model.Account
			respDataJson, _ := json.Marshal(resp.Data)
			json.Unmarshal(respDataJson, &respData)

			assert.Equal(t, cs.ExpectedResponseData, respData)

		})
	}
}

func TestBalance(t *testing.T) {
	account := &model.Account{ID: 1, Name: "test", CPF: "", Balance: *big.NewInt(100)}

	cases := map[string]struct {
		ExpectedResponseData *model.Account
		ExpectedErr          error

		Input       string
		PrepareMock func(mock *mocks.MockAccountApp)
	}{
		"return sucess": {
			Input:                "1",
			ExpectedResponseData: account,
			PrepareMock: func(mock *mocks.MockAccountApp) {
				mock.EXPECT().GetBalanceByID(gomock.Any(), 1).Return(account, nil)
			},
		},
		"return error: param ": {
			Input:       "",
			ExpectedErr: errAccountIDNotFound,
			PrepareMock: func(mock *mocks.MockAccountApp) {},
		},
		"return error: application": {
			Input:       "1",
			ExpectedErr: model.NewError(http.StatusInternalServerError, "error", nil),
			PrepareMock: func(mock *mocks.MockAccountApp) {
				mock.EXPECT().GetBalanceByID(gomock.Any(), 1).Return(nil, model.NewError(http.StatusInternalServerError, "error", nil))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)

			mockApp := mocks.NewMockAccountApp(ctrl)
			cs.PrepareMock(mockApp)

			h := handler{
				apps: &app.Container{
					Account: mockApp,
				},
			}

			req := httptest.NewRequest(http.MethodGet, "/v1/account/:id/balance", nil).WithContext(ctx)
			rec := httptest.NewRecorder()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			e := echo.New()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/account/:id/balance")
			c.SetParamNames("id")
			c.SetParamValues(cs.Input)

			err := h.balance(c)

			assert.Equal(t, cs.ExpectedErr, err)

			var resp model.Response
			json.NewDecoder(rec.Body).Decode(&resp)

			var respData *model.Account
			respDataJson, _ := json.Marshal(resp.Data)
			json.Unmarshal(respDataJson, &respData)

			assert.Equal(t, cs.ExpectedResponseData, respData)

		})
	}
}

func TestCreate(t *testing.T) {
	account := &model.Account{
		Name:   "test",
		CPF:    "12345678901",
		Secret: "secret",
	}

	cases := map[string]struct {
		ExpectedErr          error
		ExpectedResponseData *model.Account

		InputBody   func() io.Reader
		PrepareMock func(mock *mocks.MockAccountApp)
	}{
		"return sucess": {
			InputBody: func() io.Reader {
				bt, _ := json.Marshal(account)
				return bytes.NewReader(bt)
			},
			ExpectedResponseData: account,
			PrepareMock: func(mock *mocks.MockAccountApp) {
				mock.EXPECT().Create(gomock.Any(), account).Return(nil)
			},
		},
		"return error: bind": {
			InputBody: func() io.Reader {
				return strings.NewReader("invalid")
			},
			ExpectedErr: errAccountCreateBind,
			PrepareMock: func(mock *mocks.MockAccountApp) {},
		},
		"return error: application ": {
			InputBody: func() io.Reader {
				bt, _ := json.Marshal(account)
				return bytes.NewReader(bt)
			},
			ExpectedErr: model.NewError(http.StatusInternalServerError, "error", nil),
			PrepareMock: func(mock *mocks.MockAccountApp) {
				mock.EXPECT().Create(gomock.Any(), account).Return(model.NewError(http.StatusInternalServerError, "error", nil))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)

			mockApp := mocks.NewMockAccountApp(ctrl)
			cs.PrepareMock(mockApp)

			h := handler{
				apps: &app.Container{
					Account: mockApp,
				},
			}

			req := httptest.NewRequest(http.MethodPost, "/v1/account", cs.InputBody()).WithContext(ctx)
			rec := httptest.NewRecorder()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			e := echo.New()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/account")

			err := h.create(c)

			assert.Equal(t, cs.ExpectedErr, err)

			var resp model.Response
			json.NewDecoder(rec.Body).Decode(&resp)

			var respData *model.Account
			respDataJson, _ := json.Marshal(resp.Data)
			json.Unmarshal(respDataJson, &respData)

			assert.Equal(t, cs.ExpectedResponseData, respData)

		})
	}
}
