package main

import (
	"github/dmytrodemianchuk/crud_api/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handler("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
