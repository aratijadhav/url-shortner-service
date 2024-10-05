package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aratijadhav/url-shortner-service/routeutils"
	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()

	createRouteURL := routeutils.CreateRoute()

	router.HandleFunc(createRouteURL.Url, createRouteURL.HandlerName).Methods(createRouteURL.Methods...)

	fmt.Println("Starting on 8080 port")
	log.Fatal(http.ListenAndServe(":8080", router))
}
