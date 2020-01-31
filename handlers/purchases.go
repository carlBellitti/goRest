package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
)

type PurchasesResponse struct {
	Success   bool     `json:"success"`
	Purchases []string `json:"purchases"`
}

func PurchasesRouteHandler(e *Env, w http.ResponseWriter, r *http.Request) error {

	ids, ok := r.URL.Query()["id"]

	if !ok || len(ids[0]) < 1 {
		return StatusError{http.StatusBadRequest, errors.New("Bad Request")}
	}
	id := ids[0]

	if !DoesIDMatchCookie(id, r) {
		return StatusError{http.StatusUnauthorized, errors.New("Unauthorized")}
	}

	purchases := e.Storage.GetPurchases(id)
	jResponse, errr := json.Marshal(PurchasesResponse{true, purchases})

	cookieInfo := CookieInfo{ShouldSet: false, ShouldExpire: false, ID: id, SessionID: ""}
	return HandlerParseResponseAndSetCookie(w, jResponse, errr, cookieInfo)
}
