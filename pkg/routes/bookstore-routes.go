package routes

import (
	"github.com/Mr-man7352/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

// RegisterBookStoreRoutes sets up the routes for the bookstore API
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/v1/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/v1/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/v1/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/api/v1/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/v1/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

}
