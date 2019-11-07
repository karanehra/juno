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

//CreateList creates a list
func CreateList(res http.ResponseWriter, req *http.Request) {
	var list interfaces.Model = &models.List{}
	generics.CreateMethodGenericHandler(list, res, req)
}

//GetListsInBoard gets all the lists in the board
func GetListsInBoard(res http.ResponseWriter, req *http.Request) {
	boardID := mux.Vars(req)["board"]
	filter := util.CreateKeyValueFilter("boardid", boardID)
	results, err := database.FindInCollectionByFilter("lists", filter)
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	util.SendSuccessReponse(res, results)
}
