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
	MasterRouter.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	MasterRouter.HandleFunc("/user/login", controllers.AuthenticateUser).Methods("POST")
	MasterRouter.HandleFunc("/boards", controllers.CreateBoard).Methods("POST")
	MasterRouter.HandleFunc("/boards/{user}", controllers.GetUserBoards).Methods("GET")
	MasterRouter.HandleFunc("/lists", controllers.CreateList).Methods("POST")
	MasterRouter.HandleFunc("/lists/{board}", controllers.GetListsInBoard).Methods("GET")
	MasterRouter.HandleFunc("/cards", controllers.CreateCard).Methods("POST")
	MasterRouter.HandleFunc("/cards/{list}", controllers.CreateCard).Methods("GET")
	MasterRouter.HandleFunc("/articles", controllers.GetArticles).Methods("GET")
	MasterRouter.HandleFunc("/articles", controllers.PurgeArticles).Queries("page", "{page}", "size", "{size}").Methods("DELETE")
	return MasterRouter
}
