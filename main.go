package main

import (
	"fmt"
	"juno/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const PORT = 3000
	MasterRouter := mux.NewRouter()
	MasterRouter.HandleFunc("/article", controllers.CreateArticle).Methods("POST")
	MasterRouter.HandleFunc("/article", controllers.GetOne).Methods("GET")
	fmt.Printf("Server started on PORT:%d", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), MasterRouter))
}
