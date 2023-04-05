package router

import (
	"todo-react-golang/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/todos", middleware.GetAllTodos).Methods("GET")
	router.HandleFunc("/todo", middleware.CreateTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", middleware.GetTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", middleware.CompleteTodo).Methods("PUT")
	router.HandleFunc("/todo/{id}", middleware.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTodo).Methods("PUT")

	return router
}
