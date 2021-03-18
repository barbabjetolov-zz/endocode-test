package utilities

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/fatih/camelcase"
	log "github.com/sirupsen/logrus"
)

// compile time variables
var (
	GitCommit   string
	ProjectName string
)

const (
	isoFormat = "2006-01-02"
)

// struct for versionz response marshal
type ResponseVersionz struct {
	GitCommit   string `json:"git_commit"`
	ProjectName string `json:"project_name"`
}

// struct for logging
type Log struct {
	Status    int    `json:"http_status"`
	Timestamp string `json:"timestamp"`
	Request   string `json:"request"`
}

// handler for the /helloworld endpoint
func HandlerHelloworld(w http.ResponseWriter, r *http.Request) {

	var status int
	var message string

	defer LogRequest(&status, r)

	switch r.Method {
	case "GET":
		q := r.URL.Query()
		if name, isIn := q["name"]; isIn {
			status = http.StatusOK
			message = "Hello " + strings.Join(camelcase.Split(name[0]), " ")
		} else {
			status = http.StatusOK
			message = "Hello Stranger"
		}
	default:
		status = http.StatusNotFound
		message = "Method not found!"
	}

	w = writeResponse(w, message, status)
}

// handler for the /versionz endpoint
func HandlerVersionz(w http.ResponseWriter, r *http.Request) {

	var status int
	var message string

	defer LogRequest(&status, r)

	switch r.Method {
	case "GET":
		resStruct := &ResponseVersionz{GitCommit: GitCommit,
			ProjectName: ProjectName}

		resBytes, err := json.Marshal(resStruct)
		if err != nil {

			status = http.StatusInternalServerError
			message = "Internal Server Error!"

			break
		}
		status = http.StatusOK
		message = string(resBytes)
	default:
		status = http.StatusNotFound
		message = "Method not found!"
	}

	w = writeResponse(w, message, status)
}

func LogRequest(status *int, r *http.Request) {

	log.WithFields(log.Fields{
		"status": *status,
		"query":  r.URL.RawQuery,
	}).Info(r.Method + " " + r.URL.Host + r.URL.Path)
}

func writeResponse(w http.ResponseWriter, message string, status int) http.ResponseWriter {

	w.WriteHeader(status)
	w.Write([]byte(message))

	return w
}
