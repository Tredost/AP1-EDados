package produtos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func handleAdicionarProduto(w http.ResponseWriter, r *http.Request) {
	var produto Produto
	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	produto.ID = produtoID
	produtoID++
	listaProdutos.AdicionarProduto(produto)
	w.WriteHeader(http.StatusCreated)
}

func handleObterProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	produto, encontrado := listaProdutos.BuscarProdutoByID(id)
	if !encontrado {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(produto)
}

func handleRemoverProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	listaProdutos.RemoverProduto(id)
	w.WriteHeader(http.StatusNoContent)
}

func handleObterTodosProdutos(w http.ResponseWriter, r *http.Request) {
	produtos := listaProdutos.ListarProdutos()
	if len(produtos) == 0 {
		fmt.Fprintln(w, "Não há produtos cadastrados")
		return
	}
	json.NewEncoder(w).Encode(produtos)
}
