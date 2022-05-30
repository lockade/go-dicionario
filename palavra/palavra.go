package palavra

import (
	"Dicionario/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type palavra struct {
	ID        string `json:"ID"`
	Nome      string `json:"Nome"`
	Descricao string `json:"Descricao"`
}

type allPalavras []palavra

func GetAllPalavra(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusBadRequest)
	var palavras_arr allPalavras

	db := banco.OpenConnection()
	defer banco.CloseConnection(db)

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

func CreatePalavra(w http.ResponseWriter, r *http.Request) {
	var newPalavra palavra
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Dados incorretos :C")
	}

	json.Unmarshal(body, &newPalavra)

	db := banco.OpenConnection()
	defer banco.CloseConnection(db)

	insert, err := db.Query("INSERT INTO palavra (nome, descricao) VALUES (?, ?)", newPalavra.Nome, newPalavra.Descricao)
	insert.Close()
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newPalavra)
	}

}

func GetOnePalavra(w http.ResponseWriter, r *http.Request) {
	palavraID := mux.Vars(r)["id"]

	var selectedPalavra palavra

	db := banco.OpenConnection()
	defer banco.CloseConnection(db)

	err := db.QueryRow("SELECT id, nome, descricao FROM palavra WHERE id = ?", palavraID).Scan(&selectedPalavra.ID, &selectedPalavra.Nome, &selectedPalavra.Descricao)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(selectedPalavra)
	}
}

func UpdatePalavra(w http.ResponseWriter, r *http.Request) {
	palavraID := mux.Vars(r)["id"]
	var updatedPalavra palavra

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Falta Dados!")
	}

	json.Unmarshal(body, &updatedPalavra)

	db := banco.OpenConnection()
	defer banco.CloseConnection(db)

	update, err := db.Query("UPDATE palavra SET nome = ?, descricao = ? WHERE id = ?", updatedPalavra.Nome, updatedPalavra.Descricao, palavraID)
	update.Close()

	if err == nil {
		fmt.Fprintf(w, "A palavra com o id %v foi alterada", palavraID)
	}
}

func DeletePalavra(w http.ResponseWriter, r *http.Request) {
	palavraID := mux.Vars(r)["id"]

	db := banco.OpenConnection()
	defer banco.CloseConnection(db)

	update, err := db.Query("DELETE FROM palavra WHERE id = ?", palavraID)
	update.Close()

	if err == nil {
		fmt.Fprintf(w, "A palavra com o id %v foi deletada", palavraID)
	}
}
