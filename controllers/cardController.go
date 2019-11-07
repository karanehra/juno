package controllers

import (
	"juno/generics"
	"juno/interfaces"
	"juno/models"
	"net/http"
)

//CreateCard creates a new card
func CreateCard(res http.ResponseWriter, req *http.Request) {
	var card interfaces.Model = &models.Card{}
	generics.CreateMethodGenericHandler(card, res, req)
}
