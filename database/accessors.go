package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FindInCollectionByFilter returns values from the provided collection
//based on the filter
func FindInCollectionByFilter(collection string, filter primitive.D) ([]bson.M, error) {
	coll := DB.Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
