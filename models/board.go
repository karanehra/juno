package models

import (
	"context"
	"fmt"
	"juno/database"
	"juno/generics"
	"juno/util"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//Board defines the schema of a user created task board
type Board struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"userId"`
}

//Validate crosschecks a board created from a request body
func (board *Board) Validate() []string {
	return generics.GenericModelInstanceValidator(board)
}

//CreateAndSendResponse adds a userInstance to DB and sends appropriate response
func (board *Board) CreateAndSendResponse(res http.ResponseWriter) {
	coll := database.DB.Collection("boards")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := coll.InsertOne(ctx, board)
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	util.SendSuccessCreatedResponse(res, result)
}

//GetByUserIDAndSendResponse finds a user boards and writes the to the ReponseWriter provided
func (board *Board) GetByUserIDAndSendResponse(userID string, res http.ResponseWriter) {
	collection := database.DB.Collection("boards")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, util.CreateKeyValueFilter("userId", userID))
	if err != nil {
		fmt.Println("errored")
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
		fmt.Println("append")
	}
	fmt.Println(len(results))
	util.SendSuccessReponse(res, results)
	if err := cur.Err(); err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
}
