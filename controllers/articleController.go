package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"juno/database"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//GetArticles fetches articles from db
func GetArticles(res http.ResponseWriter, req *http.Request) {
	coll := database.DB.Collection("articles")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	results, err := coll.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(results)
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(results)
}
