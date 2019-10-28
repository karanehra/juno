package controllers

import (
	"context"
	"encoding/json"
	"juno/database"
	"juno/models"
	"juno/util"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CreateUser handles the user creation endpoint
func CreateUser(res http.ResponseWriter, req *http.Request) {
	var userInstance models.User
	json.NewDecoder(req.Body).Decode(&userInstance)
	if err := userInstance.Validate(); len(err) > 0 {
		util.SendBadRequestResponse(res, err)
		return
	}
	userInstance.Password = util.CreateHashSHA(userInstance.Password)
	coll := database.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := coll.InsertOne(ctx, userInstance)
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	util.SendSuccessCreatedResponse(res, result)
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
	coll := database.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "email", Value: credentials.Email}}
	err := coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	if !user.IsPasswordCorrect(credentials.Password) {
		resBody := make(map[string]interface{})
		resBody["message"] = "Incorrect Credentials"
		util.SendUnauthorizedResponse(res, resBody)
		return
	}
	user.Password = ""
	util.SendSuccessReponse(res, user)
}
