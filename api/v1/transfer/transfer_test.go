package transfer

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/patrickchagastavares/conta-corrent/app"
	"github.com/patrickchagastavares/conta-corrent/mocks"
	"github.com/patrickchagastavares/conta-corrent/model"
	"github.com/patrickchagastavares/conta-corrent/test"
	"github.com/patrickchagastavares/conta-corrent/utils/session"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	lists := []*model.Transfer{
		{ID: 1, OriginID: 1, DestinationID: 2, Amount: *big.NewInt(100)},
		{ID: 2, OriginID: 1, DestinationID: 2, Amount: *big.NewInt(100)},
	}
	sessionjwt := session.SessionJWT{
		AccountOriginID: 1,
		Name:            "test",
		StandardClaims:  jwt.StandardClaims{},
	}

	cases := map[string]struct {
		ExpectedResponseData []*model.Transfer
		ExpectedErr          error

		session     session.SessionJWT
		PrepareMock func(mock *mocks.MockTransferApp)
	}{
		"return sucess": {
			session:              sessionjwt,
			ExpectedResponseData: lists,
			PrepareMock: func(mock *mocks.MockTransferApp) {
				mock.EXPECT().ListByID(gomock.Any(), 1).Return(lists, nil)
			},
		},
		"return error: application ": {
			session:     sessionjwt,
			ExpectedErr: model.NewError(http.StatusInternalServerError, "error", nil),
			PrepareMock: func(mock *mocks.MockTransferApp) {
				mock.EXPECT().ListByID(gomock.Any(), 1).Return(nil, model.NewError(http.StatusInternalServerError, "error", nil))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)

			ctx = context.WithValue(ctx, "session", cs.session)

			mockApp := mocks.NewMockTransferApp(ctrl)
			cs.PrepareMock(mockApp)

			h := handler{
				apps: &app.Container{
					Transfer: mockApp,
				},
			}

			req := httptest.NewRequest(http.MethodGet, "/v1/transfer", nil).WithContext(ctx)
			rec := httptest.NewRecorder()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			e := echo.New()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/transfer")

			err := h.list(c)

			assert.Equal(t, cs.ExpectedErr, err)

			var resp model.Response
			json.NewDecoder(rec.Body).Decode(&resp)

			var respData []*model.Transfer
			respDataJson, _ := json.Marshal(resp.Data)
			json.Unmarshal(respDataJson, &respData)

			assert.Equal(t, cs.ExpectedResponseData, respData)

		})
	}
}

func TestCreate(t *testing.T) {
	sessionjwt := session.SessionJWT{
		AccountOriginID: 1,
		Name:            "test",
		StandardClaims:  jwt.StandardClaims{},
	}
	transfer := &model.Transfer{
		OriginID:      1,
		DestinationID: 2,
		Amount:        *big.NewInt(100),
	}

	cases := map[string]struct {
		ExpectedErr error

		session     session.SessionJWT
		InputBody   func() io.Reader
		PrepareMock func(mock *mocks.MockTransferApp)
	}{
		"return sucess": {
			session: sessionjwt,
			InputBody: func() io.Reader {
				bt, _ := json.Marshal(transfer)
				return bytes.NewReader(bt)
			},
			PrepareMock: func(mock *mocks.MockTransferApp) {
				mock.EXPECT().Create(gomock.Any(), transfer).Return(nil)
			},
		},
		"return error: bind": {
			session: sessionjwt,
			InputBody: func() io.Reader {
				return strings.NewReader("invalid")
			},
			ExpectedErr: errTransferBind,
			PrepareMock: func(mock *mocks.MockTransferApp) {},
		},
		"return error: application ": {
			session: sessionjwt,
			InputBody: func() io.Reader {
				bt, _ := json.Marshal(transfer)
				return bytes.NewReader(bt)
			},
			ExpectedErr: model.NewError(http.StatusInternalServerError, "error", nil),
			PrepareMock: func(mock *mocks.MockTransferApp) {
				mock.EXPECT().Create(gomock.Any(), transfer).Return(model.NewError(http.StatusInternalServerError, "error", nil))
			},
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			ctrl, ctx := test.NewController(t)

			ctx = context.WithValue(ctx, "session", cs.session)

			mockApp := mocks.NewMockTransferApp(ctrl)
			cs.PrepareMock(mockApp)

			h := handler{
				apps: &app.Container{
					Transfer: mockApp,
				},
			}

			req := httptest.NewRequest(http.MethodPost, "/v1/transfer", cs.InputBody()).WithContext(ctx)
			rec := httptest.NewRecorder()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			e := echo.New()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/transfer")

			err := h.create(c)

			assert.Equal(t, cs.ExpectedErr, err)

		})
	}
}
