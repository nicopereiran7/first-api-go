package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicopereiran7/first-api-go/common"
	"github.com/nicopereiran7/first-api-go/routes"
)

func main() {
	common.Migrate()

	router := mux.NewRouter()
	routes.SetUserRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose en el puerto 9000")
	log.Println(server.ListenAndServe())
}
