package web

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller struct { //aarmazena as pessoas
	nextId uint
	store  map[uint]Person //dicionario que associa um id com uma pessoa
}
type Person struct {
	Id        uint   `json:id`
	Name      string `json:name`
	Height    string `json:height`
	Mass      string `json:mass`
	HairColor string `json:hair-color`
	SkinColor string `json:skin-color`
	EyeColor  string `json:eye-color`
	BirthYear string `json: birth-year`
	Gender    string `json:gender`
}

func NewController() *Controller {
	return &Controller{
		nextId: 1,
		store:  make(map[uint]Person),
	}

}
func (c *Controller) HandlePing(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong!")

}

func (c *Controller) HandleCreatePerson(w http.ResponseWriter, r *http.Request) {

	p := Person{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}
	id := c.nextId
	p.Id = id
	c.store[id] = p
	c.nextId++

	w.WriteHeader(http.StatusNoContent)
}

func (c *Controller) HandleListPeople(w http.ResponseWriter, r *http.Request) {
	people := make([]Person, len(c.store))
	i := 0
	for _, value := range c.store {
		people[i] = value
		i++
	}
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode((people))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Controller) HandleGetPersonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personIdRaw := vars["id"]
	strconv.ParseUint(personIdRaw, 10, 32)
	personId, err := strconv.Atoi(personIdRaw)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uint(personId)
	p, found := c.store[id]
	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode((p))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}