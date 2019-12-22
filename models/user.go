package models

import (
	"context"
	"crypto/sha1"
	"juno/database"
	"juno/generics"
	"juno/util"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//User defines the user mongoDB schema
type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

//UserCredentials defines the login api schema
type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Validate ensures that user is in a correct format
func (user *User) Validate() []string {
	return generics.GenericModelInstanceValidator(user)
}

//Validate checks login request payload
func (userCredentials *UserCredentials) Validate() []string {
	var errorData []string = []string{}
	if userCredentials.Email == "" {
		errorData = append(errorData, "email: field is required")
	}
	if userCredentials.Password == "" {
		errorData = append(errorData, "password: field is required")
	}
	return errorData
}

//IsPasswordCorrect hashes the password and checks if it matches the password hash
func (user *User) IsPasswordCorrect(password string) bool {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	res := hasher.Sum(nil)
	if string(res) == user.Password {
		return true
	}
	return false
}

//CreateAndSendResponse adds a userInstance to DB and send appropriate response
func (user *User) CreateAndSendResponse(res http.ResponseWriter) {
	var userInstance User
	userInstance.Password = util.GetSHAToHex(userInstance.Password)
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

//FindOne finds and returns a mongo doc
func (user *User) FindOne(filter primitive.D) *mongo.SingleResult {
	coll := database.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return coll.FindOne(ctx, filter)
}
