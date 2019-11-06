package generics

import (
	"encoding/json"
	"juno/interfaces"
	"juno/util"
	"net/http"
)

//CreateMethodGeneric tries to implement a model agnostic CREATE method
func CreateMethodGeneric(
	model interfaces.Model,
	res http.ResponseWriter,
	req *http.Request) {
	json.NewDecoder(req.Body).Decode(model)
	if err := model.Validate(); len(err) > 0 {
		util.SendBadRequestResponse(res, err)
		return
	}
	model.CreateAndSendResponse(res)
}
