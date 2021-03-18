package main

import (
	"net/http"
	"strings"

	"github.com/fatih/camelcase"
	logrus "github.com/sirupsen/logrus"
)

var gitHash string

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
// func HandlerVersionz(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "GET":
// 		res := &utilities.ResponseVersionz{GitHash: gitHash,
// 			ProjectName: projectName}
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"git_hash": "get called 2"}`))
// 	default:
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte(`Method not found!`))
// 	}
// }

func main() {
	http.HandleFunc("/helloworld", HandlerHelloworld)
	// http.HandleFunc("/versionz", HandlerVersionz)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}
