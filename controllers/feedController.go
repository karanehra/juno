package controllers

import (
	"context"
	"juno/database"
	"juno/generics"
	"juno/interfaces"
	"juno/models"
	"juno/util"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

//CreateFeed handles the feed post endpoint
func CreateFeed(res http.ResponseWriter, req *http.Request) {
	var feed interfaces.Model = &models.Feed{}
	generics.CreateMethodGenericHandler(feed, res, req)
}

//GetFeeds handles the feed getter route
func GetFeeds(res http.ResponseWriter, req *http.Request) {
	data, err := database.FindInCollectionByFilter("feeds", bson.D{})
	if err != nil {
		util.SendServerErrorResponse(res, "An error occured while querying")
		return
	}
	util.SendSuccessReponse(res, data)
}

//GetFeedByID handles the singular feed getter route
func GetFeedByID(res http.ResponseWriter, req *http.Request) {
	feedID := mux.Vars(req)["feedID"]
	data, err := database.FindInCollectionByID("feeds", feedID)
	if err != nil {
		util.SendServerErrorResponse(res, "An error occured while querying")
		return
	}
	util.SendSuccessReponse(res, data)
}

//UpdateFeedByID handles the singular feed updater route
func UpdateFeedByID(res http.ResponseWriter, req *http.Request) {
	util.SendSuccessReponse(res, "")
}

//PurgeFeeds handles the delete all feeds route
func PurgeFeeds(res http.ResponseWriter, req *http.Request) {
	coll := database.DB.Collection("feeds")
	coll.DeleteMany(context.TODO(), bson.D{})
	util.SendSuccessReponse(res, map[string]string{})
}
