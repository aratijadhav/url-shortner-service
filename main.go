package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()

	fmt.Println("Starting on 8080 port")
	log.Fatal(http.ListenAndServe(":8080", router))
}
