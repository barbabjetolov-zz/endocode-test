package utilities

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

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
		t.Errorf("got %q, want 200", code)
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
		t.Errorf("got %q, want 200", code)
	}
}

func TestEmptyQuery(t *testing.T) {

	want := "Invalid Request!"

	request, err := http.NewRequest(http.MethodGet, "/helloworld", nil)
	if err != nil {
		t.Error(err)
	}
	response := httptest.NewRecorder()

	q := request.URL.Query()
	q.Add("name", "")
	request.URL.RawQuery = q.Encode()

	HandlerHelloworld(response, request)

	got := response.Body.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	code := response.Code

	if code != 400 {
		t.Errorf("got %q, want 400", code)
	}
}

func TestFuzzQueryParam(t *testing.T) {

	// this test will be repeated 50 times
	count := 50

	for i := 0; i < count; i++ {
		want := "Invalid Request!"

		request, err := http.NewRequest(http.MethodGet, "/helloworld", nil)
		if err != nil {
			t.Error(err)
		}
		response := httptest.NewRecorder()

		q := request.URL.Query()
		q.Add(randString(10), randString(10))
		request.URL.RawQuery = q.Encode()

		HandlerHelloworld(response, request)

		got := response.Body.String()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		code := response.Code

		if code != 400 {
			t.Errorf("got %q, want 400", code)
		}
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
		t.Errorf("got %q, want 200", code)
	}

	if gotHeader != wantHeader {
		t.Errorf("got header %q, want %q", gotHeader, wantHeader)
	}
}

func TestWriteResponse(t *testing.T) {

	status := http.StatusOK
	message := "Test message!"

	response := httptest.NewRecorder()

	writeResponse(&status, &message, response)

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	got := string(bodyBytes)

	if got != message {
		t.Errorf("got %q, want %q", got, message)
	}

	code := response.Result().StatusCode

	if code != status {
		t.Errorf("got %q, want %q", code, status)
	}

}

// helpers
func randString(length int) string {

	rand.Seed(time.Now().UnixNano())

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}
