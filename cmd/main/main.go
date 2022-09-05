package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rizky-putra19/books-project-go/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	fmt.Println("server running on localhost:3001")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3001", r))
}
