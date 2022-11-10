package main

import (
	"connectDB/server"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", server.SearchUSers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.SearchUSer).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", server.DeletUser).Methods(http.MethodDelete)

	fmt.Println("Rodando na porta :5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
