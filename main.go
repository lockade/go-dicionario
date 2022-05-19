package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type palavra struct {
	ID        string `json:"ID"`
	Nome      string `json:"Nome"`
	Descricao string `json:"Descricao"`
}

type allPalavras []palavra

var palavras = allPalavras{
	{
		ID:        "1",
		Nome:      "Abacate",
		Descricao: "É um fruta doce! (as vezes né)",
	},
}

func createPalavra(w http.ResponseWriter, r *http.Request) {
	var newPalavra palavra
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Dados incorretos :C")
	}

	json.Unmarshal(body, &newPalavra)

	//add in array palavras
	palavras = append(palavras, newPalavra)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newPalavra)
}

func getOnePalavra(w http.ResponseWriter, r *http.Request) {
	palavraID := mux.Vars(r)["id"]

	for _, singlePalavra := range palavras {
		if singlePalavra.ID == palavraID {
			json.NewEncoder(w).Encode(singlePalavra)
		}
	}
}

func updatePalavra(w http.ResponseWriter, r *http.Request) {
	palavraID := mux.Vars(r)["id"]
	var updatedPalavra palavra

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Falta Dados!")
	}

	json.Unmarshal(body, &updatedPalavra)

	for i, singlePalavra := range palavras {
		if singlePalavra.ID == palavraID {
			singlePalavra.Nome = updatedPalavra.Nome
			singlePalavra.Descricao = updatedPalavra.Descricao
			palavras = append(palavras[:i], singlePalavra)
			json.NewEncoder(w).Encode(singlePalavra)
		}
	}

}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá!!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/palavra", createPalavra).Methods("POST")
	router.HandleFunc("/palavra/{id}", getOnePalavra).Methods("GET")
	router.HandleFunc("/palavra/{id}", updatePalavra).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":9090", router))
}
