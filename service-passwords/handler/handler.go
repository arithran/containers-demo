package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

func Hash(w http.ResponseWriter, r *http.Request) {
	// 1. Get password from URL
	pass := mux.Vars(r)["pass"]

	// 2. Hash Password
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		sendJSONResponse(w, newResponse("error generating password"), http.StatusBadRequest)
		return
	}

	// 3. Write Response (with hashed password)
	sendJSONResponse(w, newResponse(string(bytes)), http.StatusOK)
}

func Check(w http.ResponseWriter, r *http.Request) {
	// 1. Get Payload (Password + Hash)
	pass := mux.Vars(r)["pass"]
	hash := mux.Vars(r)["hash"]

	// 2. Compare
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		// Write Failed Response
		sendJSONResponse(w, newResponse("invalid"), http.StatusUnauthorized)
		return
	}

	// 4. Write Successful Response
	sendJSONResponse(w, newResponse("valid"), http.StatusOK)
}

type Response struct {
	Message string `json:"message"`
}

func newResponse(msg string) Response {
	return Response{msg}
}

func sendJSONResponse(w http.ResponseWriter, resp interface{}, code int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Error("sendJSONResponse() error while encoding. err = ", err)
	}
}
