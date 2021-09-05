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

type SWPeopleRepository struct{}

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
		return nil, err
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

func (r SWPeopleRepository) Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	res, err := Collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return errors.New("Could not delete")
	}
	fmt.Println(res.DeletedCount)
	return nil
}

func (r SWPeopleRepository) Update(id string, person domain.Person) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	params := map[string]string{"name": person.Name, "height": person.Height, "mass": person.Mass, "haircolor": person.HairColor, "skincolor": person.SkinColor, "eyecolor": person.EyeColor, "gender": person.Gender, "films": string(person.Films)}
	for k, p := range params {
		if p == "" {
			continue
		}
		fmt.Println(k, p)
		res, err := Collection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.D{{"$set", bson.D{{k, p}}}})
		if err != nil {
			return errors.New("Could not update field!")
		}
		fmt.Println(res)
	}
	return nil

}
