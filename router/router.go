package router

import (
	"juno/controllers"
	"juno/middlewares"
	"os"

	"github.com/gorilla/mux"
)

//SetupRouter creates app route definitions
func SetupRouter() *mux.Router {
	MasterRouter := mux.NewRouter()
	MasterRouter.Use(middlewares.CORSMiddleware)
	MasterRouter.Use(middlewares.JSONContentMiddleware)
	MasterRouter.Use(middlewares.LoggerMiddleware)
	if os.Getenv("APP_AUTH") == "on" {
		MasterRouter.Use(middlewares.AuthMiddleware)
	}
	MasterRouter.HandleFunc("/test", controllers.TestController).Methods("GET")
	MasterRouter.HandleFunc("/user", controllers.CreateUser).Methods("OPTIONS", "POST")
	MasterRouter.HandleFunc("/user/login", controllers.AuthenticateUser).Methods("OPTIONS", "POST")
	MasterRouter.HandleFunc("/boards", controllers.CreateBoard).Methods("POST")
	MasterRouter.HandleFunc("/boards/{user}", controllers.GetUserBoards).Methods("GET")
	MasterRouter.HandleFunc("/lists", controllers.CreateList).Methods("POST")
	MasterRouter.HandleFunc("/lists/{board}", controllers.GetListsInBoard).Methods("GET")
	MasterRouter.HandleFunc("/cards", controllers.CreateCard).Methods("POST")
	MasterRouter.HandleFunc("/cards/{list}", controllers.CreateCard).Methods("GET")

	MasterRouter.
		HandleFunc("/articles", controllers.GetArticles).
		Methods("GET", "OPTIONS").
		Queries("page", "{page}", "size", "{size}", "query", "{query}")

	MasterRouter.HandleFunc("/tags", controllers.GetTags).Methods("GET")

	MasterRouter.HandleFunc("/articles", controllers.PurgeArticles).Methods("OPTIONS", "DELETE")

	MasterRouter.HandleFunc("/feeds", controllers.GetFeeds).Methods("OPTIONS", "GET")
	MasterRouter.HandleFunc("/feeds/{feedID}", controllers.UpdateFeedByID).Methods("OPTIONS", "POST")
	MasterRouter.HandleFunc("/feeds", controllers.CreateFeed).Methods("OPTIONS", "POST")
	MasterRouter.HandleFunc("/feeds", controllers.PurgeFeeds).Methods("OPTIONS", "DELETE")
	MasterRouter.HandleFunc("/feeds/{feedID}", controllers.GetFeedByID).Methods("GET")

	MasterRouter.HandleFunc("/process", controllers.GetProcessesHandler).Methods("GET")
	MasterRouter.HandleFunc("/process", controllers.CreateProcessHandler).Methods("OPTIONS", "POST")

	MasterRouter.HandleFunc("/dataset", controllers.GetDataset).Methods("OPTIONS", "GET")
	return MasterRouter
}
