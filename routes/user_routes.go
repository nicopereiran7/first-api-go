package routes

import (
	"github.com/gorilla/mux"
	"github.com/nicopereiran7/first-api-go/controllers"
)

func SetUserRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/user").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/find/{id}", controllers.Get).Methods("GET")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete", controllers.Delete).Methods("DELETE")
}
