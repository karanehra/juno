package router

import (
	"juno/controllers"

	"github.com/gorilla/mux"
)

//MasterRouter contains route definitions for the app
var MasterRouter *mux.Router

func init() {
	MasterRouter := mux.NewRouter()
	MasterRouter.HandleFunc("/article", controllers.CreateArticle).Methods("POST")
	MasterRouter.HandleFunc("/article", controllers.GetOne).Methods("GET")
}
