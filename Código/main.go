package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type QuantidadeProduto struct {
	ID         int `json:"id"`
	Quantidade int `json:"quantidade"`
}

var (
	listaProdutos ListaProdutos
	filaPedidos   FilaPedidos
	produtoID     int = 1
	pedidoID      int = 1
	métricas      Métricas
	lojaAberta    bool
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/produto", handleAdicionarProduto).Methods("POST")
	r.HandleFunc("/produto/{id}", handleObterProduto).Methods("GET")
	r.HandleFunc("/produto/{id}", handleRemoverProduto).Methods("DELETE")
	r.HandleFunc("/produtos", handleObterTodosProdutos).Methods("GET")
	r.HandleFunc("/pedido", handleIncluirPedido).Methods("POST")
	r.HandleFunc("/pedidos", handleObterPedidosAtivos).Methods("GET")
	r.HandleFunc("/abrir", handleAbrirLoja).Methods("POST")
	r.HandleFunc("/fechar", handleFecharLoja).Methods("POST")
	r.HandleFunc("/metricas", handleObterMétricas).Methods("GET")
	http.ListenAndServe(":8080", r)
}
