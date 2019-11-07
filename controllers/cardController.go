package controllers

import (
	"juno/database"
	"juno/generics"
	"juno/interfaces"
	"juno/models"
	"juno/util"
	"net/http"

	"github.com/gorilla/mux"
)

//CreateCard creates a new card
func CreateCard(res http.ResponseWriter, req *http.Request) {
	var card interfaces.Model = &models.Card{}
	generics.CreateMethodGenericHandler(card, res, req)
}

//GetCardInList creates a new card
func GetCardInList(res http.ResponseWriter, req *http.Request) {
	listID := mux.Vars(req)["list"]
	filter := util.CreateKeyValueFilter("listid", listID)
	results, err := database.FindInCollectionByFilter("cards", filter)
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	util.SendSuccessReponse(res, results)
}
