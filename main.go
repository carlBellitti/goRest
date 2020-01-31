package main

import (
	"log"
	"net/http"

	"test.com/goRest/handlers"
	"test.com/goRest/storage"
)

func main() {

	// Initialise our app-wide environment with the services/info we need.
	env := &handlers.Env{}
	s := storage.StorageRepo{}
	s.SetMockData()
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Note that we're using http.Handle, not http.HandleFunc. The
	// latter only accepts the http.HandlerFunc type, which is not
	// what we have here.
	http.Handle("/login", handlers.Handler{env, handlers.LoginRouteHandler, http.MethodPost, false})
	http.Handle("/purchases", handlers.Handler{env, handlers.PurchasesRouteHandler, http.MethodGet, true})
	http.Handle("/auth", handlers.Handler{env, handlers.AuthRouteHandler, http.MethodGet, true})
	http.Handle("/logout", handlers.Handler{env, handlers.LogoutRouteHandler, http.MethodPost, true})

	// Logs the error if ListenAndServe fails.
	log.Fatal(http.ListenAndServe(":8000", nil))
}
