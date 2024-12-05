package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    LoginDataResponse `json:"data"`
}

type LoginDataResponse struct {
	MovieData []string `json:"moviedata"`
	ID        string   `json:"id"`
	FullName  string   `json:"fullname"`
	UserName  string   `json:"username"`
	Password  string   `json:"password"`
}

func LoginRouteHandler(e *Env, w http.ResponseWriter, r *http.Request) error {

	var l LoginRequest
	test := "1234"

	test2, _ := strconv.ParseInt(test, 10, 64)
	log.Println(uint32(test2))

	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		return StatusError{http.StatusBadRequest, errors.New("Bad Request")}
	}

	var response LoginResponse
	id := ""
	sessionID := ""
	shouldSetCookie := false
	if l.UserName == "" || l.Password == "" {
		return StatusError{http.StatusUnauthorized, errors.New("Unauthorized")}
	}
	log.Println("LOGIN " + l.UserName + ", " + l.Password)
	data, err := e.Storage.CheckLogin(l.UserName, l.Password)
	if err != nil {
		log.Println("UNAUTHORIZED " + err.Error())
		return StatusError{http.StatusUnauthorized, err}
	}
	response = LoginResponse{true, "User Found", LoginDataResponse{data.MovieData, data.ID, data.FullName, data.UserName, data.Password}}
	id = data.ID
	sessionID = data.SessionID
	shouldSetCookie = true
	jResponse, errr := json.Marshal(response)
	cookieInfo := CookieInfo{ShouldSet: shouldSetCookie, ShouldExpire: false, ID: id, SessionID: sessionID}
	return HandlerParseResponseAndSetCookie(w, jResponse, errr, cookieInfo)
}
