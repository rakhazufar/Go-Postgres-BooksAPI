package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rakhazufar/go-postgres/pkg/routes"
	_ "gorm.io/driver/postgres"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStore(r)
	// http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}