package main

import (
	"fmt"
	"juno/router"
	"log"
	"net/http"
)

func main() {
	const PORT = 3000
	fmt.Printf("Server started on PORT:%d\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), router.SetupRouter()))
}
