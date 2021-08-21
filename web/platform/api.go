package platform

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/lumacielz/star-wars-api/domain"
)

const baseUrl = "https://swapi.dev/api/people"

type personFilms struct {
	Films []string `json:"films"`
}

type apiPeopleResponse struct {
	Results []personFilms
}

func GetPersonFilms(p domain.Person) (int, error) {
	search := url.QueryEscape(p.Name)
	targetUrl := fmt.Sprintf("%s?search=%s", baseUrl, search)

	resp, err := http.Get(targetUrl)
	if err != nil {
		return 0, err
	}

	var peopleResp apiPeopleResponse
	err = json.NewDecoder(resp.Body).Decode(&peopleResp)
	if err != nil {
		return 0, err
	}

	if len(peopleResp.Results) == 0 {
		return 0, fmt.Errorf("empty star wars API response")
	}

	person := peopleResp.Results[0]
	return len(person.Films), nil
}
