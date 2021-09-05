package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lumacielz/star-wars-api/database"
	"github.com/lumacielz/star-wars-api/web"
)

func main() {
	repo := database.SWPeopleRepository{}
	ctrl := web.NewController(repo)
	r := mux.NewRouter()
	r.Methods(http.MethodPost).Path("/person").HandlerFunc(ctrl.HandleCreatePerson)
	r.Methods(http.MethodGet).Path("/person").HandlerFunc(ctrl.HandleListPeople)
	r.Methods(http.MethodGet).Path("/person/{id}").HandlerFunc(ctrl.HandleGetPersonById)
	r.Methods(http.MethodDelete).Path("/person/{id}").HandlerFunc(ctrl.HandleDeletePerson)
	r.Methods(http.MethodPut).Path("/person/{id}").HandlerFunc(ctrl.HandleUpdatePerson)

	log.Println("starting server...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Printf("[ERROR] %s\n", err.Error())
		os.Exit(1)
	}

}
