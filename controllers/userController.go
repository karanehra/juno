package controllers

import (
	"encoding/json"
	"fmt"
	"juno/database"
	"juno/generics"
	"juno/models"
	"juno/util"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

//CreateUser handles the user creation endpoint
func CreateUser(res http.ResponseWriter, req *http.Request) {
	var user *models.User = &models.User{}
	if req.Body == nil {
		util.SendBadRequestResponse(res, map[string]interface{}{"errors": "Invalid Request"})
		return
	}
	json.NewDecoder(req.Body).Decode(user)
	email := user.Email
	filter := util.CreateKeyValueFilter("email", email)
	users, err := database.FindInCollectionByFilter("users", filter)
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	if len(users) > 1 {
		util.SendBadRequestResponse(res, map[string]interface{}{"errors": "Email ID is already in user"})
		return
	}
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
	token, err := util.CreateJWTForIssuer(user.Email)
	fmt.Println(err)
	payload := map[string]interface{}{
		"user":  user,
		"token": token,
	}
	util.SendSuccessReponse(res, payload)
}

//TestController used for testing and hacking
func TestController(res http.ResponseWriter, req *http.Request) {
	results, _ := database.FindInCollectionByFilter("users", bson.D{})
	util.SendSuccessReponse(res, results)
}
