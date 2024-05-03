package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/produto", produtos.AdicionarProduto).Methods("POST")
	r.HandleFunc("/produto/{id}", produtos.ObterProduto).Methods("GET")
	r.HandleFunc("/produto/{id}", produtos.RemoverProduto).Methods("DELETE")
	r.HandleFunc("/produtos", produtos.ObterTodosProdutos).Methods("GET")
	r.HandleFunc("/pedido", pedidos.IncluirPedido).Methods("POST")
	r.HandleFunc("/pedidos", pedidos.ObterPedidosAtivos).Methods("GET")
	r.HandleFunc("/abrir", loja.AbrirLoja).Methods("POST")
	r.HandleFunc("/fechar", loja.FecharLoja).Methods("POST")
	r.HandleFunc("/metricas", metricas.ObterMétricas).Methods("GET")
	http.ListenAndServe(":8080", r)
}
