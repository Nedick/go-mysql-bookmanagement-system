package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nedick/go-mysql-bookmanagement-system/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server on port 7171")
	log.Fatal(http.ListenAndServe("localhost:7171", r))
}
