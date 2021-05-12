package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

var servicePasswordAddress = ""

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var db = map[string]User{}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// 1.Decode Payload
	u := User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		sendJSONResponse(w, newResponse("error decoding body"), http.StatusBadRequest)
		return
	}

	// 2.External Call to Passwords Service
	resp, err := http.Get(fmt.Sprintf("%s/passwords/hash/%s", servicePasswordAddress, u.Password))
	if err != nil {
		sendJSONResponse(w, newResponse("password service is unavailable"), http.StatusServiceUnavailable)
		return
	}
	if resp.StatusCode != http.StatusOK {
		sendJSONResponse(w, newResponse("error status "+resp.Status), http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()
	out := Response{}
	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		logrus.Error(err)
	}
	u.Password = out.Message

	// 3.Save to Database
	db[u.Email] = u

	// 4. Write Response
	sendJSONResponse(w, u, http.StatusCreated)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	// 1. Select all from Database
	res := make([]User, 0, len(db))
	for _, u := range db {
		res = append(res, u)
	}

	// 2. Write Response
	sendJSONResponse(w, res, http.StatusOK)
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

func init() {
	servicePasswordAddress = os.Getenv("SERVICE_PASSWORD_ADDRESS")
	if servicePasswordAddress == "" {
		logrus.Error("servicePasswordAddress is not set in ENV")

	}
}
