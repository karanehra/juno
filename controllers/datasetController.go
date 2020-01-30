package controllers

import (
	"juno/database"
	"juno/util"
	"net/http"
)

//GetDataset handles the dataset supplier endpoint
func GetDataset(res http.ResponseWriter, req *http.Request) {
	articleCount, err := database.GetCollectionDocumentCount("articles")
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}

	feedCount, feedErr := database.GetCollectionDocumentCount("feeds")
	if feedErr != nil {
		util.SendServerErrorResponse(res, feedErr.Error())
		return
	}

	payload := map[string]interface{}{
		"feedCount":    feedCount,
		"articleCount": articleCount,
	}

	util.SendSuccessReponse(res, payload)
}
