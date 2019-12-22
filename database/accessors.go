package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

//GetPaginatedArticles is used to paginate article results
func GetPaginatedArticles(pageSize int32, pageNo int64) ([]bson.M, error) {
	coll := DB.Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	skip := ((pageNo - 1) * int64(pageSize))
	options := options.Find().SetLimit(int64(pageSize)).SetSkip(skip)
	cur, err := coll.Find(ctx, bson.M{}, options)
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
