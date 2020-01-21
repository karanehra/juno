package controllers

import (
	"juno/util"
	"net/http"
)

//GetProcessesHandler handles the process listing endpoint
func GetProcessesHandler(res http.ResponseWriter, req *http.Request) {
	util.SendSuccessReponse(res, "")
}

//CreateProcessHandler handlers the process creation endpoint
func CreateProcessHandler(res http.ResponseWriter, req *http.Request) {
	util.SendSuccessCreatedResponse(res, "")
}
