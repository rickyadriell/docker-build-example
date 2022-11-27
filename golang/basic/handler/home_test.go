package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil) // create a request
	response := httptest.NewRecorder()                       // create a response recorder

	t.Run("Home handler", func(t *testing.T) {
		HomeHandler(response, request) // call the handler

		got := response.Body.String() // get the response
		want := "Hello World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
