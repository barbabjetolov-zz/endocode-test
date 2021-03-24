package utilities

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fatih/camelcase"
)

func TestHellostranger(t *testing.T) {

	want := "Hello Stranger"

	request, err := http.NewRequest(http.MethodGet, "/helloworld", nil)
	if err != nil {
		t.Error(err)
	}
	response := httptest.NewRecorder()

	HandlerHelloworld(response, request)

	got := response.Body.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	code := response.Code

	if code != 200 {
		t.Errorf("got %q, want 200", got)
	}
}

func TestHelloname(t *testing.T) {

	name := "EdoardoRizzardi"
	want := "Hello " + strings.Join(camelcase.Split(name), " ")

	request, err := http.NewRequest(http.MethodGet, "/helloworld", nil)
	if err != nil {
		t.Error(err)
	}
	response := httptest.NewRecorder()

	q := request.URL.Query()
	q.Add("name", name)
	request.URL.RawQuery = q.Encode()

	HandlerHelloworld(response, request)

	got := response.Body.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	code := response.Code

	if code != 200 {
		t.Errorf("got %q, want 200", got)
	}
}

func TestVersionz(t *testing.T) {

	want := "{\"git_commit\":\"" + GitCommit + "\",\"project_name\":\"" + ProjectName + "\"}"
	wantHeader := "application/json; charset=rtf-8"

	request, err := http.NewRequest(http.MethodGet, "/versionz", nil)
	if err != nil {
		t.Error(err)
	}
	response := httptest.NewRecorder()

	HandlerVersionz(response, request)

	got := response.Body.String()
	gotHeader := response.Header().Get("Content-type")

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	code := response.Code

	if code != 200 {
		t.Errorf("got %q, want 200", got)
	}

	if gotHeader != wantHeader {
		t.Errorf("got header %q, want %q", gotHeader, wantHeader)
	}
}
