package controllers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
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

//TestController used for testing and hacking
func TestController(res http.ResponseWriter, req *http.Request) {
	randomizer := rand.Reader
	key, err := rsa.GenerateKey(randomizer, 1024)
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	str := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	secretMessage := []byte("abcdefghijklmnopqrstuvwxyz")
	encMessage, _ := rsa.EncryptPKCS1v15(randomizer, &key.PublicKey, secretMessage)
	fmt.Println(hex.EncodeToString(encMessage))
	util.SendSuccessReponse(res, map[string]string{"publicKey": hex.EncodeToString(str), "l": string(len(str))})
}
