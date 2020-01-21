package controllers

import (
	"encoding/json"
	"juno/models"
	"juno/util"
	"net/http"

	"github.com/karanehra/schemas"
)

//GetProcessesHandler handles the process listing endpoint
func GetProcessesHandler(res http.ResponseWriter, req *http.Request) {
	data, err := models.GetAllProcesses()
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	util.SendSuccessReponse(res, data)
}

//CreateProcessHandler handlers the process creation endpoint
func CreateProcessHandler(res http.ResponseWriter, req *http.Request) {
	var process *schemas.Process = &schemas.Process{Status: "CREATED"}
	json.NewDecoder(req.Body).Decode(process)
	data, err := models.CreateProcess(*process)
	if err != nil {
		util.SendServerErrorResponse(res, err.Error())
		return
	}
	util.SendSuccessCreatedResponse(res, data)
}
