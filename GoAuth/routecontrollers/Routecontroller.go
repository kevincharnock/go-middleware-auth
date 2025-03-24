package routecontrollers

import (
	"oauth/handlers"
	middleware "oauth/middlewares"

	"github.com/gorilla/mux"
)

func Rotinhadesgracada() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/protected", middleware.TokenVerifyMiddleware(handlers.ProtectedHandler)).Methods("GET")

	return router
}
