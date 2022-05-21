package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

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

func getAllPalavra(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusBadRequest)
	var palavras_arr allPalavras

	results, err := db.Query("SELECT id, nome, descricao FROM palavra")
	if err == nil {
		for results.Next() {
			var p palavra
			err = results.Scan(&p.ID, &p.Nome, &p.Descricao)
			if err == nil {
				palavras_arr = append(palavras_arr, p)
			}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(palavras_arr)
	}
}

func createPalavra(w http.ResponseWriter, r *http.Request) {
	var newPalavra palavra
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Dados incorretos :C")
	}

	json.Unmarshal(body, &newPalavra)

	//add in array palavras
	insert, err := db.Query("INSERT INTO palavra (nome, descricao) VALUES (?, ?)", newPalavra.Nome, newPalavra.Descricao)
	insert.Close()
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newPalavra)
	}
	// palavras = append(palavras, newPalavra)

}

func getOnePalavra(w http.ResponseWriter, r *http.Request) {
	palavraID := mux.Vars(r)["id"]

	var selectedPalavra palavra

	err := db.QueryRow("SELECT id, nome, descricao FROM palavra WHERE id = ?", palavraID).Scan(&selectedPalavra.ID, &selectedPalavra.Nome, &selectedPalavra.Descricao)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(selectedPalavra)
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

	update, err := db.Query("UPDATE palavra SET nome = ?, descricao = ? WHERE id = ?", updatedPalavra.Nome, updatedPalavra.Descricao, palavraID)
	update.Close()

	if err == nil {
		fmt.Fprintf(w, "A palavra com o id %v foi alterada", palavraID)
	}
}

func deletePalavra(w http.ResponseWriter, r *http.Request) {
	palavraID := mux.Vars(r)["id"]

	update, err := db.Query("DELETE FROM palavra WHERE id = ?", palavraID)
	update.Close()

	if err == nil {
		fmt.Fprintf(w, "A palavra com o id %v foi deletada", palavraID)
	}
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá!!")
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/dicionario")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/palavra", getAllPalavra).Methods("GET")
	router.HandleFunc("/palavra", createPalavra).Methods("POST")
	router.HandleFunc("/palavra/{id}", getOnePalavra).Methods("GET")
	router.HandleFunc("/palavra/{id}", updatePalavra).Methods("PATCH")
	router.HandleFunc("/palavra/{id}", deletePalavra).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9090", router))
}
