package routes

import (
	"api-practice1/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {

	router.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/api/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/books", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT")

}
