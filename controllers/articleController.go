package controllers

import (
	"context"
	"juno/database"
	"juno/util"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//GetOne is a test getter function
func GetOne(res http.ResponseWriter, req *http.Request) {
	collection := database.DB.Collection("articles")
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	var results []bson.M
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			util.SendServerErrorResponse(res, err.Error())
			return
		}
		results = append(results, result)
	}
	util.SendSuccessReponse(res, results)
	if err := cur.Err(); err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
}
