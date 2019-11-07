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

//CreateBoard creates a board
func CreateBoard(res http.ResponseWriter, req *http.Request) {
	var board interfaces.Model = &models.Board{}
	generics.CreateMethodGenericHandler(board, res, req)
}

//GetUserBoards gets the boards created againt the provided userID
func GetUserBoards(res http.ResponseWriter, req *http.Request) {
	userID := mux.Vars(req)["user"]
	filter := util.CreateKeyValueFilter("userid", userID)
	results, err := database.FindInCollectionByFilter("boards", filter)
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	util.SendSuccessReponse(res, results)
}
