package models

import (
	"context"
	"juno/database"
	"juno/generics"
	"juno/util"
	"net/http"
	"time"
)

//List type defines the list schema
type List struct {
	Title   string `json:"title"`
	BoardID string `json:"boardid"`
	UserID  string `json:"userid"`
}

//Validate method crosschecks and validates the list values
func (list *List) Validate() []string {
	return generics.GenericModelInstanceValidator(list)
}

//CreateAndSendResponse method creates a list in database from a
//validated list
func (list *List) CreateAndSendResponse(res http.ResponseWriter) {
	coll := database.DB.Collection("lists")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := coll.InsertOne(ctx, list)
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	util.SendSuccessCreatedResponse(res, result)
}
