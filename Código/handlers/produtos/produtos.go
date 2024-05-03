package IANZINHO

import (
	p "IANZINHO/processamento"
	"net/http"
)

func HandleAdicionarProduto(w http.ResponseWriter, r *http.Request) {
	p.AdicionarProduto(w, r)
}

func handleObterProduto(w http.ResponseWriter, r *http.Request) {
	p.ObterProduto(w, r)
}

func handleObterTodosProdutos(w http.ResponseWriter, r *http.Request) {
	p.ObterTodosProdutos(w, r)
}
