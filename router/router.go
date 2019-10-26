package router

import (
	"juno/controllers"
	"juno/middlewares"

	"github.com/gorilla/mux"
)

//SetupRouter creates app route definitions
func SetupRouter() *mux.Router {
	MasterRouter := mux.NewRouter()
	MasterRouter.Use(middlewares.LoggerMiddleware)
	MasterRouter.Use(middlewares.JSONContentMiddleware)
	MasterRouter.HandleFunc("/article", controllers.CreateArticle).Methods("POST")
	MasterRouter.HandleFunc("/article", controllers.GetOne).Methods("GET")
	return MasterRouter
}
