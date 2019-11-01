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
	MasterRouter.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	MasterRouter.HandleFunc("/user/login", controllers.AuthenticateUser).Methods("POST")
	MasterRouter.HandleFunc("/boards", controllers.CreateBoard).Methods("POST")
	MasterRouter.HandleFunc("/boards/{user}", controllers.GetUserBoards).Methods("GET")
	return MasterRouter
}
