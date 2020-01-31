package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type AuthResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    AuthDataResponse `json:"data"`
}

type AuthDataResponse struct {
	MovieData []string `json:"moviedata"`
	ID        string   `json:"id"`
	FullName  string   `json:"fullname"`
	UserName  string   `json:"username"`
	Password  string   `json:"password"`
}

func AuthRouteHandler(e *Env, w http.ResponseWriter, r *http.Request) error {

	var response AuthResponse

	shouldSetCookie := false
	cookieData, err := IsCookieAuthenticated(r, e.Storage)
	if err != nil {
		//return StatusError{http.StatusUnauthorized, errors.New("Unauthorized")}
		response = AuthResponse{false, "Unauthorized", AuthDataResponse{}}
	}
	log.Println("AUTH Id: " + cookieData.ID)
	data, err := e.Storage.GetAppDataById(cookieData.ID)
	if err != nil {
		log.Println("SERVER ERROR " + err.Error())
		return StatusError{http.StatusInternalServerError, err}
	}
	response = AuthResponse{true, "User Found", AuthDataResponse{data.MovieData, data.ID, data.FullName, data.UserName, data.Password}}
	shouldSetCookie = true
	jResponse, errr := json.Marshal(response)
	cookieInfo := CookieInfo{ShouldSet: shouldSetCookie, ShouldExpire: false, ID: data.ID, SessionID: data.SessionID}
	return HandlerParseResponseAndSetCookie(w, jResponse, errr, cookieInfo)
}
