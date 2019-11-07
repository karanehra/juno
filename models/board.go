package models

import (
	"context"
	"juno/database"
	"juno/generics"
	"juno/util"
	"net/http"
	"time"
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
