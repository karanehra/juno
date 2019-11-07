package models

import (
	"context"
	"juno/database"
	"juno/generics"
	"juno/util"
	"net/http"
	"time"
)

//Card type defines the schema for a card
type Card struct {
	Description string   `json:"description"`
	BoardID     string   `json:"boardid"`
	MemberIDs   []string `json:"members"`
}

//Validate crosschecks a board created from a request body
func (card *Card) Validate() []string {
	return generics.GenericModelInstanceValidator(card)
}

//CreateAndSendResponse adds a userInstance to DB and sends appropriate response
func (card *Card) CreateAndSendResponse(res http.ResponseWriter) {
	coll := database.DB.Collection("cards")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := coll.InsertOne(ctx, card)
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	util.SendSuccessCreatedResponse(res, result)
}
