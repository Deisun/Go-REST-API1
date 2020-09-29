package controllers

import (
	"api-practice1/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// DUMMY DATA
var books = []models.Book{
	{
		ID: "1",
		Title: "AmazeBalls",
		Author: []models.Author{
			{"Rob", "Douma"}, {"Tom", "Mitcham"},
		},
	},
	{
		ID: "2",
		Title: "SpaceBalls",
		Author: []models.Author{{"Tania", "Ocasio"}},
	},
	{
		ID: "3",
		Title: "Nunez",
		Author: []models.Author{{"Shelley", "Zimmerman"}},
	},
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	var newBook models.Book

	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not Decode new book"))
		return
	}

	books = append(books, newBook)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newBook); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not Encode book"))
		return
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not convert ID to integer"))
		return
	}

	if id <= 0 || id > len(books) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is out of range"))
		return
	}

	books := append(books[:id - 1], books[id:]...)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("There was a problem encoding"))
		return
	}

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("there was a problem"))
		return
	}

	if id > len(books) || id < 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Book Id does not exist"))
		return
	}

	book := books[id - 1]

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("There was a problem Encoding"))
		return
	}
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("There was a problem Encoding"))
		return
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	paramID := mux.Vars(r)["id"]
	id, err := strconv.Atoi(paramID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("There was a problem converting ID to str"))
		return
	}

	if id >= len(books) || id < 1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No book found with that ID."))
		return
	}

	var updatedBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		log.Println(err)
		return
	}

	books[id - 1] = updatedBook

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedBook); err != nil {
		log.Println(err)
		return
	}
}
