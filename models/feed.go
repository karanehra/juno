package models

import (
	"context"
	"juno/database"
	"juno/generics"
	"juno/util"
	"net/http"
	"time"
)

//Feed defines the model for a feed
type Feed struct {
	URL   string   `json:"url"`
	Title string   `json:"title"`
	Tags  []string `json:"tags"`
}

func (feed *Feed) Validate() []string {
	return generics.GenericModelInstanceValidator(feed)
}

//CreateAndSendResponse adds a userInstance to DB and sends appropriate response
func (feed *Feed) CreateAndSendResponse(res http.ResponseWriter) {
	coll := database.DB.Collection("feeds")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := coll.InsertOne(ctx, feed)
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	util.SendSuccessCreatedResponse(res, result)
}
