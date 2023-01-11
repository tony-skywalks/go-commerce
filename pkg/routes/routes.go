package routes

import (
	"github.com/gorilla/mux"
	"github.com/tony-skywalks/my-web/pkg/controllers"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/user", controllers.Accounts).Methods("GET", "PUT", "POST")
	router.HandleFunc("/user/login", controllers.Login).Methods("POST", "GET")
}
