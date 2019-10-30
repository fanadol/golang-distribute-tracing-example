package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fanadol/golang-distribute-tracing-example/server"
	"github.com/gorilla/mux"
)

func main() {

	serverRepo := server.NewDatabase()
	serverService := server.NewServerService(serverRepo)
	serverHandler := server.NewServerHandler(serverService)

	router := mux.NewRouter()

	router.HandleFunc("/post", serverHandler.Create).Methods("POST")
	router.HandleFunc("/post", serverHandler.Get).Methods("GET")

	fmt.Println("starting web server at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", router))
}
