package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lumacielz/star-wars-api/web"
)

func main() {
	ctrl := web.NewController()
	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/ping").HandlerFunc(ctrl.HandlePing)
	r.Methods(http.MethodPost).Path("/person").HandlerFunc(ctrl.HandleCreatePerson)
	r.Methods(http.MethodGet).Path("/person").HandlerFunc(ctrl.HandleListPeople)
	log.Println("starting server...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Printf("[ERROR] %s\n", err.Error())
		os.Exit(1)
	}
}