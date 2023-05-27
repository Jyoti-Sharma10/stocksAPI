package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	router "stocksAPI/routers"
)

	func main() {
		r := router.Router()
		fmt.Println("Starting server on the port 8080...")
		log.Fatal(http.ListenAndServe(":8080", r))
	}

