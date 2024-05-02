package produto

import (
	"encoding/json"
	"net/http"
)

type Produto struct {
	ID        int     `json:"id"`
	Nome      string  `json:"nome"`
	Descrição string  `json:"descrição"`
	Valor     float64 `json:"valor"`
}

func AdicionarProduto(w http.ResponseWriter, r *http.Request) {
	var Produto Produto
	err := json.NewDecoder(r.Body).Decode(&Produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ProdutoID := Produto.ID
	ProdutoID++
	ListaProdutos.AdicionarProduto(Produto)
	w.WriteHeader(http.StatusCreated)
}
