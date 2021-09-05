package domain

import (
	"errors"
)

var ErrNotFound = errors.New("person not found")

type People []Person

type Person struct {
	Name      string `bson:name`
	Height    string `bson:height`
	Mass      string `bson:mass`
	HairColor string `bson:hair-color`
	SkinColor string `bson:skin-color`
	EyeColor  string `bson:eye-color`
	BirthYear string `bson: birth-year`
	Gender    string `bson:gender`
	Films     int    `bson:"films"`
}
