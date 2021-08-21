package domain

import "errors"

var ErrNotFound = errors.New("person not found")

type People []Person

type Person struct {
	Id        uint64 `json:id`
	Name      string `json:name`
	Height    string `json:height`
	Mass      string `json:mass`
	HairColor string `json:hair-color`
	SkinColor string `json:skin-color`
	EyeColor  string `json:eye-color`
	BirthYear string `json: birth-year`
	Gender    string `json:gender`
	Films     int    `json:"films"`
}
