package web

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lumacielz/star-wars-api/database/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type PeopleRepository interface {
	Create(p domain.Person) error
	List() ([]bson.M, error)
	GetById(id string) ([]bson.M, error)
	Update(id string, person domain.Person) error
	Delete(id string) error
}
type Controller struct {
	repo PeopleRepository
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
	people, err := c.repo.List()
	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(people)
	w.Write(data)
}

func (c *Controller) HandleGetPersonById(w http.ResponseWriter, r *http.Request) {
	personId, err := c.parsePersonId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	res, err := c.repo.GetById(personId)
	if err == domain.ErrNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	//err = json.NewEncoder(w).Encode((person))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(res)
	w.Write(data)

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

func (c *Controller) parsePersonId(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	personIdRaw := vars["id"]
	personId := string(personIdRaw)
	return personId, nil
}
func (c *Controller) parsePersonBody(r *http.Request) (domain.Person, error) {
	var p domain.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		return domain.Person{}, err
	}

	return p, err
}
