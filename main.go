package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/barbabjetolov/endocode-test/http-service/pkg/utilities"
	log "github.com/sirupsen/logrus"
)

func init() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {

	port, z := os.LookupEnv("LISTENING_PORT")

	if z {
		if _, err := strconv.Atoi(port); err != nil {
			log.Fatal(err, ": Listening port must be a number")
		}
	}
	if port == "" {
		port = "8080"
	}

	log.Info("Listening on port " + port)

	http.HandleFunc("/helloworld", utilities.HandlerHelloworld)
	http.HandleFunc("/versionz", utilities.HandlerVersionz)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
