package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"juno/database"
	"juno/models"
	"juno/util"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

//CreateBoard creates a board
func CreateBoard(res http.ResponseWriter, req *http.Request) {
	var boardInstance models.Board
	json.NewDecoder(req.Body).Decode(&boardInstance)
	fmt.Println(boardInstance)
	if err := boardInstance.Validate(); len(err) > 0 {
		util.SendBadRequestResponse(res, err)
		return
	}
	boardInstance.CreateAndSendResponse(res)
}

//GetUserBoards gets the boards created againt the provided userID
func GetUserBoards(res http.ResponseWriter, req *http.Request) {
	userID := mux.Vars(req)["user"]
	fmt.Println(userID)
	collection := database.DB.Collection("boards")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, util.CreateKeyValueFilter("userid", userID))
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
	fmt.Println(len(results))
	util.SendSuccessReponse(res, results)
	if err := cur.Err(); err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
}
