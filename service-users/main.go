package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/arithran/containers-demo/service-users/handler"
	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const port = 8000

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users/create", handler.CreateUser).Methods("POST")
	r.HandleFunc("/users/list", handler.ListUsers).Methods("GET")

	// start server
	logrus.Infof("Starting up on %d...", port)
	listenAddr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(listenAddr, gh.LoggingHandler(os.Stdout, r))
	if err != nil {
		logrus.Error(err)
	}
}
