package controllers

import (
	"context"
	"encoding/json"
	"juno/database"
	"juno/models"
	"juno/util"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//CreateArticle adds an article to database
func CreateArticle(res http.ResponseWriter, req *http.Request) {
	var v models.Article
	json.NewDecoder(req.Body).Decode(&v)
	if err := v.Validate(); len(err) > 0 {
		util.SendBadRequestResponse(res, err)
		return
	}
	coll := database.DB.Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := coll.InsertOne(ctx, v)
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	util.SendSuccessCreatedResponse(res, result)
}

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
