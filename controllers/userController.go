package controllers

import (
	"encoding/json"
	"juno/generics"
	"juno/interfaces"
	"juno/models"
	"juno/util"
	"net/http"
)

//CreateUser handles the user creation endpoint
func CreateUser(res http.ResponseWriter, req *http.Request) {
	var user interfaces.Model = &models.User{}
	generics.CreateMethodGenericHandler(user, res, req)
}

//AuthenticateUser crosschecks user credentials
func AuthenticateUser(res http.ResponseWriter, req *http.Request) {
	var credentials models.UserCredentials
	var user models.User
	json.NewDecoder(req.Body).Decode(&credentials)
	if err := credentials.Validate(); len(err) > 0 {
		util.SendBadRequestResponse(res, err)
		return
	}
	filter := util.CreateKeyValueFilter("email", credentials.Email)
	err := user.FindOne(filter).Decode(&user)
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	if !user.IsPasswordCorrect(credentials.Password) {
		util.SendUnauthorizedResponse(res, "Incorrect Credentials")
		return
	}
	user.Password = ""
	util.SendSuccessReponse(res, user)
}
