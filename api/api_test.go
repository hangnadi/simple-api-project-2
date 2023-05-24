package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/hangnadi/simple-api-project-2/api"
)

func TestWriteResponseJSON(t *testing.T) {

	w := httptest.NewRecorder()

	data := []byte{}

	WriteResponseJSON(w, data, http.StatusOK, nil)
}

func TestWriteResponseJSONWithHeader(t *testing.T) {

	w := httptest.NewRecorder()

	data := []byte{}
	headers := map[string]string{}
	headers["foo"] = "bar"

	WriteResponseJSON(w, data, http.StatusOK, headers)
}
