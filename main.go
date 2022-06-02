package main

import (
	"Dicionario/palavra"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/palavra", palavra.GetAllPalavra).Methods("GET")
	router.HandleFunc("/palavra", palavra.CreatePalavra).Methods("POST")
	router.HandleFunc("/palavra/{id}", palavra.GetOnePalavra).Methods("GET")
	router.HandleFunc("/palavra/{id}", palavra.UpdatePalavra).Methods("PATCH")
	router.HandleFunc("/palavra/{id}", palavra.DeletePalavra).Methods("DELETE")

	// corsObj := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":9090", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
