package util

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CreateKeyValueFilter returns a filter object for finding mongo documents
func CreateKeyValueFilter(key string, value string) bson.D {
	return bson.D{primitive.E{Key: key, Value: value}}
}
