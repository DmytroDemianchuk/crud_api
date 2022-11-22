package routes

import (
	"github.com/dmytrodemianchuk/crud-api/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandlerFunc("/book/", controllers.CreateBook).Methhods("POST")
	router.HandlerFunc("/book", controllers.GetBook).Methhods("GET")
	router.HandlerFunc("/book/{bookid}", controllers.GetBookid).Methhods("GET")
	router.HandlerFunc("/book/{bookid", controllers.UpdateBook).Methhods("PUT")
	router.HandlerFunc("/book{bookid", controllers.DeleteBook).Methhods("DELETE")
}
