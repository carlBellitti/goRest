package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
)

type LogoutRequest struct {
	Id string `json:"id"`
}

type LogoutResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func LogoutRouteHandler(e *Env, w http.ResponseWriter, r *http.Request) error {

	var l LogoutRequest
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		return StatusError{http.StatusBadRequest, errors.New("Bad Request")}
	}

	id := l.Id

	e.Storage.Logout(id)
	jResponse, errr := json.Marshal(LogoutResponse{true, "SUCCESS"})

	cookieInfo := CookieInfo{ShouldSet: true, ShouldExpire: true, ID: "", SessionID: ""}
	return HandlerParseResponseAndSetCookie(w, jResponse, errr, cookieInfo)
}
