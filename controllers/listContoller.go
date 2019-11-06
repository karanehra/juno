package controllers

import (
	"juno/generics"
	"juno/interfaces"
	"juno/models"
	"net/http"
)

//CreateList creates a list
func CreateList(res http.ResponseWriter, req *http.Request) {
	var list interfaces.Model = &models.List{}
	generics.CreateMethodGenericHandler(list, res, req)
}
