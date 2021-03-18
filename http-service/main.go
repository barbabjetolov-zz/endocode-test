package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/barbabjetolov/endocode-test/http-service/pkg/utilities"
	"github.com/fatih/camelcase"
	logrus "github.com/sirupsen/logrus"
)

// compile time variables
var (
	GitCommit   string
	ProjectName string
)

// handler for the /helloworld endpoint
func HandlerHelloworld(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		q := r.URL.Query()
		if name, isIn := q["name"]; isIn {

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`Hello ` + strings.Join(camelcase.Split(name[0]), " ")))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`Hello Stranger`))
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`Method not found!`))
	}
}

// handler for the /versionz endpoint
func HandlerVersionz(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		resStruct := &utilities.ResponseVersionz{GitCommit: GitCommit,
			ProjectName: ProjectName}

		resBytes, err := json.Marshal(resStruct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`Internal server error!`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resBytes))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`Method not found!`))
	}
}

func main() {
	http.HandleFunc("/helloworld", HandlerHelloworld)
	http.HandleFunc("/versionz", HandlerVersionz)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}
