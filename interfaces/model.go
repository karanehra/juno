package interfaces

import "net/http"

//Model defines the model interfaces
type Model interface {
	Validate() []string
	CreateAndSendResponse(res http.ResponseWriter)
}
