package main

import (
	"net/http"
	"os"

	"github.com/barbabjetolov/endocode-test/http-service/pkg/utilities"
	log "github.com/sirupsen/logrus"
)

func init() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {

	http.HandleFunc("/helloworld", utilities.HandlerHelloworld)
	http.HandleFunc("/versionz", utilities.HandlerVersionz)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
