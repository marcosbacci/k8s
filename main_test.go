package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreeting(t *testing.T) {
	t.Run("Greeting", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		greeting(response, request)

		got := response.Body.String()
		want := "<b>Code.education Rocks!</b>"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
