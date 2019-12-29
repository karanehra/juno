package main

import (
	"fmt"
	"juno/router"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {
	const PORT = 3007
	fmt.Printf("Server started on PORT:%d\n", PORT)
	log.Fatal(
		http.ListenAndServe(fmt.Sprintf(":%d", PORT),
			handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}))(router.SetupRouter())))
}
