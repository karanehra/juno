package generics

import (
	"juno/interfaces"
	"juno/util"
	"net/http"
)

//CreateMethodGenericHandler implements a model agnostic CREATE method
func CreateMethodGenericHandler(
	model interfaces.Model,
	res http.ResponseWriter,
	req *http.Request) {
	if err := model.Validate(); len(err) > 0 {
		responseBody := map[string]interface{}{"validationErrors": err}
		util.SendBadRequestResponse(res, responseBody)
		return
	}
	model.CreateAndSendResponse(res)
}
