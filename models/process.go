package models

import (
	"context"
	"juno/database"

	"go.mongodb.org/mongo-driver/bson"
)

//Process defines a task to be sent to a processor
type Process struct {
	Name   string `json:"processName"`
	Status string `json:"status"`
}

//GetAllProcesses fetches all processes from the database
func GetAllProcesses() ([]bson.M, error) {
	coll := database.DB.Collection("process")
	cur, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var results []bson.M
	err = cur.All(context.TODO(), results)

	return results, err

}
