package generics

import (
	"encoding/json"
	"juno/interfaces"
	"juno/util"
	"net/http"
)

//CreateMethodGenericHandler implements a model agnostic CREATE method
func CreateMethodGenericHandler(
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
