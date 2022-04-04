package model

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	cases := map[string]struct {
		ExpectedErr string

		InputError Error
	}{
		"return: valid with sucess": {
			InputError:  Error{HTTPCode: http.StatusInternalServerError, Message: "test", Detail: nil},
			ExpectedErr: "code: 500 - message: test - detail: <nil>",
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			err := cs.InputError.Error()

			assert.Equal(t, err, cs.ExpectedErr)

		})
	}
}

func TestGetHTTPCode(t *testing.T) {
	cases := map[string]struct {
		ExpectedHTTPCode int

		InputError error
	}{
		"return: valid with sucess": {
			InputError:       &Error{HTTPCode: http.StatusBadRequest, Message: "test", Detail: nil},
			ExpectedHTTPCode: http.StatusBadRequest,
		},
		"return error": {
			InputError:       nil,
			ExpectedHTTPCode: http.StatusInternalServerError,
		},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			httpCode := GetHTTPCode(cs.InputError)

			assert.Equal(t, httpCode, cs.ExpectedHTTPCode)

		})
	}
}
