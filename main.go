package main

import (
	"Dicionario/palavra"
	"log"
	"net/http"

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

	log.Fatal(http.ListenAndServe(":9090", router))
}
