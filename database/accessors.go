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
	cursorError := cur.All(context.TODO(), &results)
	if cursorError != nil {
		return nil, cursorError
	}
	return results, nil
}

//FindInCollectionByID finds a document in collection with the given ID
func FindInCollectionByID(collection string, ID string) (bson.D, error) {
	coll := DB.Collection(collection)
	result := coll.FindOne(context.TODO(), bson.M{"_id": ID})
	if result.Err() != nil {
		return nil, result.Err()
	}
	doc := bson.D{}
	err := result.Decode(doc)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

//GetPaginatedArticles is used to paginate article results
func GetPaginatedArticles(pageSize int32, pageNo int64, query string) ([]bson.M, error) {
	coll := DB.Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	skip := ((pageNo - 1) * int64(pageSize))
	options := options.Find().SetLimit(int64(pageSize)).SetSkip(skip)
	filter := bson.D{
		bson.E{
			Key: "description", Value: bson.D{{"$regex", primitive.Regex{Pattern: query, Options: "i"}}}}}
	cur, err := coll.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	cursorError := cur.All(context.TODO(), &results)
	if cursorError != nil {
		return nil, cursorError
	}
	return results, nil
}
