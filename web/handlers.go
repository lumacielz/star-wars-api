package web

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lumacielz/star-wars-api/domain"
)

type PeopleRepository interface {
	Create(p domain.Person) error
	List() domain.People
	Get(id uint64) (domain.Person, error)
	Update(id uint64, person domain.Person) error
	Delete(id uint64) error
}
type Controller struct { //aarmazena as pessoas
	repo PeopleRepository //dicionario que associa um id com uma pessoa
}

func NewController(r PeopleRepository) *Controller {
	return &Controller{repo: r}

}

func (c *Controller) HandleCreatePerson(w http.ResponseWriter, r *http.Request) {
	p, err := c.parsePersonBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}
	err = c.repo.Create(p)
	if err != nil {
		log.Printf("[ERROR] failed to create person on database: %s\n", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) HandleListPeople(w http.ResponseWriter, r *http.Request) {
	people := c.repo.List()
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode((people))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Controller) HandleGetPersonById(w http.ResponseWriter, r *http.Request) {
	personId, err := c.parsePersonId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	person, err := c.repo.Get(personId)
	if err == domain.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode((person))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (c *Controller) HandleDeletePerson(w http.ResponseWriter, r *http.Request) {
	id, err := c.parsePersonId(r)
	if err != nil {
		log.Printf("[ERROR] failed to parser person id: %s\n", err)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.repo.Delete(id)
	if err == domain.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}

func (c *Controller) HandleUpdatePerson(w http.ResponseWriter, r *http.Request) {
	id, err := c.parsePersonId(r)
	if err != nil {
		log.Printf("[ERROR] failed to parser person id: %s\n", err)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	person, err := c.parsePersonBody(r)
	if err != nil {
		log.Printf("[ERROR] failed to read update person body: %s\n", err)

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.repo.Update(id, person)
	if err == domain.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *Controller) parsePersonId(r *http.Request) (uint64, error) {
	vars := mux.Vars(r)
	personIdRaw := vars["id"]
	personId, err := strconv.Atoi(personIdRaw)
	if err != nil {
		return 0, err
	}

	if personId <= 0 {
		return 0, fmt.Errorf("person id should not be zero or less")
	}

	return uint64(personId), nil
}
func (c *Controller) parsePersonBody(r *http.Request) (domain.Person, error) {
	var p domain.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		return domain.Person{}, err
	}

	return p, err
}
