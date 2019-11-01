package models

import (
	"context"
	"juno/database"
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
	var errorData []string = []string{}
	if board.Title == "" {
		errorData = append(errorData, "title: field is required")
	}
	if board.Description == "" {
		errorData = append(errorData, "description: field is required")
	}
	if board.UserID == "" {
		errorData = append(errorData, "userId: field is required")
	}

	return errorData
}

//CreateAndSendResponse adds a userInstance to DB and send appropriate response
func (board *Board) CreateAndSendResponse(res http.ResponseWriter) {
	var boardInstance Board
	coll := database.DB.Collection("boards")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := coll.InsertOne(ctx, boardInstance)
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	util.SendSuccessCreatedResponse(res, result)
}
