package produtos

import (
	"net/http"
	//IMPORTAR PROCESSAMENTO
)

func HandleAdicionarProduto(w http.ResponseWriter, r *http.Request) {
	AdicionarProduto(w, r)
}

func handleObterProduto(w http.ResponseWriter, r *http.Request) {
	ObterProduto(w, r)
}

func handleObterTodosProdutos(w http.ResponseWriter, r *http.Request) {
	ObterTodosProdutos(w, r)
}
