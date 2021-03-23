package utilities

import (
	"encoding/json"
	"net/http"
	"regexp"
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

	defer LogRequest(&status, &message, r)
	defer writeResponse(&status, &message, w)

	switch r.Method {
	case "GET":
		q := r.URL.Query()
		if len(q) > 1 {
			status = http.StatusBadRequest
			message = "Invalid request!"
		} else if len(q) == 1 {
			if name, isIn := q["name"]; isIn {

				if !isWord(name[0]) {
					status = http.StatusBadRequest
					message = "Invalid Request!"

					break
				}

				status = http.StatusOK
				message = "Hello " + strings.Join(camelcase.Split(name[0]), " ")
			} else {
				status = http.StatusBadRequest
				message = "Invalid request!"
			}
		} else if len(q) == 0 {
			status = http.StatusOK
			message = "Hello Stranger"
		}
	default:
		status = http.StatusNotFound
		message = "Method not found!"
	}
}

// handler for the /versionz endpoint
func HandlerVersionz(w http.ResponseWriter, r *http.Request) {

	var status int
	var message string

	defer LogRequest(&status, &message, r)
	defer writeResponse(&status, &message, w)

	q := r.URL.Query()

	switch r.Method {
	case "GET":
		if len(q) > 100 {
			status = http.StatusBadRequest
			message = "Invalid request! Query string too long!"
		} else if len(q) != 0 {
			status = http.StatusBadRequest
			message = "Invalid request! Empty query!"
		} else {
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
		}
	default:
		status = http.StatusNotFound
		message = "Method not found!"
	}
}

// function that logs every served request
func LogRequest(status *int, message *string, r *http.Request) {

	log.WithFields(log.Fields{
		"status":  *status,
		"query":   r.URL.RawQuery,
		"request": r.Method + " " + r.URL.Host + r.URL.Path,
	}).Info(*message)
}

func writeResponse(status *int, message *string, w http.ResponseWriter) {

	w.WriteHeader(*status)
	w.Write([]byte(*message))
}

func isWord(s string) bool {

	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(s)
}
