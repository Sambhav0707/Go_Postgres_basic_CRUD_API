package main

import (
	"go_postgres/router"
	"log"
	"net/http"
)

func main() {

	r := router.Router()

	log.Fatal(http.ListenAndServe(":8081", r))

}
