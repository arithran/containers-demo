package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/arithran/containers-demo/service-passwords/handler"
	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const port = 8001

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/passwords/hash/{pass}", handler.Hash).Methods("GET")
	r.HandleFunc("/passwords/check/{pass}/{hash}", handler.Check).Methods("GET")

	// start server
	logrus.Infof("Starting up on %d...", port)
	listenAddr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(listenAddr, gh.LoggingHandler(os.Stdout, r))
	if err != nil {
		logrus.Error(err)
	}
}
