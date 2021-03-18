package utilities

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHellostranger(t *testing.T) {

	want := "Hello Stranger"

	request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
	response := httptest.NewRecorder()

	HandlerHelloworld(response, request)

	got := response.Body.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
