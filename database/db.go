package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/lumacielz/star-wars-api/database/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection = CreateConection()

type SWPeopleRepository struct {
}

func (r SWPeopleRepository) Create(p domain.Person) error {

	res, err := Collection.InsertOne(context.Background(), p)
	if err != nil {
		return errors.New("não inserido")
	}
	fmt.Println(res.InsertedID)
	return nil
}

func (r SWPeopleRepository) List() ([]bson.M, error) {
	var res []bson.M
	cursor, err := Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("Unexpected Error")
	}
	if err = cursor.All(context.Background(), &res); err != nil {
		return nil, errors.New("Decoding Error")
	}

	return res, nil
}

func (r SWPeopleRepository) GetById(id string) ([]bson.M, error) {
	var res []bson.M
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
	}
	cursor, err := Collection.Find(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return nil, errors.New("Unexpected Error")
	}
	if err = cursor.All(context.Background(), &res); err != nil {
		return nil, errors.New("Decoding Error")
	}

	return res, nil

}
