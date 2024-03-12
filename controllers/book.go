package controllers

import (
	"encoding/json"
	"go-crud/entities"
	"go-crud/repos"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var bookRepo = repos.NewBookRepo()

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book entities.Book

	json.NewDecoder(r.Body).Decode(&book)
	book = bookRepo.Create(book)

	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusCreated)
}

func Show(w http.ResponseWriter, r *http.Request) {
	bookIdLong, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Do not understand {id}")
		return
	}

	book, err := bookRepo.GetOne(uint(bookIdLong))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Book not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
	w.WriteHeader(http.StatusOK)
}
