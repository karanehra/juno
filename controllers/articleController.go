package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"juno/database"
	"juno/models"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//CreateArticle adds an article to database
func CreateArticle(res http.ResponseWriter, req *http.Request) {
	var v models.Article
	json.NewDecoder(req.Body).Decode(&v)
	fmt.Println(v)
	coll := database.DB.Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, _ := coll.InsertOne(ctx, v)
	json.NewEncoder(res).Encode(result)
}

//GetOne is a test getter function
func GetOne(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	collection := database.DB.Collection("articles")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var results []bson.M
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
		// do something with result....
	}
	json.NewEncoder(res).Encode(results)
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
