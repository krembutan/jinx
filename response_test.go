package jinx

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupResponse() Response {
	return Response{
		Pagination: map[string]interface{}{},
	}
}

func TestSemanticVersion(t *testing.T) {
	response := setupResponse()
	version := "1.0.0"
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	w.WriteHeader(http.StatusOK) // set header code
	if got, want := w.Code, http.StatusOK; got != want {
		t.Fatalf("status code got: %d, want %d", got, want)
	}

	SemanticVersion(r, version)
	expected := map[string]interface{}{
		"foo": "bar",
	}

	err = ResponseJSONPayload(w, r, http.StatusOK, expected)
	assert.NoError(t, err)

	bytes, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	var actual Response
	err = json.Unmarshal(bytes, &actual)
	assert.NoError(t, err)

	response.Version = map[string]interface{}{
		"number": version,
	}
	response.Data = expected
	response.Meta = map[string]interface{}{"code": http.StatusText(http.StatusOK)}
	assert.Equal(t, response, actual)
}

func TestHttpErrStatus(t *testing.T) {
	status := []struct {
		label string
		code  int
		err   func(w http.ResponseWriter, r *http.Request, err error) error
	}{
		{http.StatusText(http.StatusBadRequest), http.StatusBadRequest, ErrBadRequest},
		{http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized, ErrUnauthorized},
		{http.StatusText(http.StatusForbidden), http.StatusForbidden, ErrForbidden},
		{http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed, ErrMethodNotAllowed},
		{http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError, ErrInternalServerError},
		{http.StatusText(http.StatusBadGateway), http.StatusBadGateway, ErrBadGateway},
		{http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable, ErrServiceUnavailable},
	}

	t.Run("http status", func(t *testing.T) {
		for _, tt := range status {
			t.Run(tt.label, func(t *testing.T) {
				r, err := http.NewRequest(http.MethodGet, "/", nil)
				assert.NoError(t, err)
				w := httptest.NewRecorder()
				errReq := tt.err(w, r, fmt.Errorf("%s or %v", tt.label, "constraint unique key duplicate"))
				assert.Error(t, errReq)
				if got, want := w.Code, tt.code; got != want {
					t.Fatalf("status code got: %d, want %d", got, want)
				}
			})
		}
	})
}
