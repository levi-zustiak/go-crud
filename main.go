package main

import (
	"fmt"
	"go-crud/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	RegisterBookRoutes(r)

	http.ListenAndServe(":8080", r)
}

func RegisterBookRoutes(r *mux.Router) {
	muxBase := "/api/books"

	r.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), controllers.Show).Methods("GET")
	r.HandleFunc(muxBase, controllers.Create).Methods("POST")
}
