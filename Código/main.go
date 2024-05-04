package main

import (
	"net/http"

	p "IANZINHO/processamento"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/produto", p.AdicionarProduto).Methods("POST")
	r.HandleFunc("/produto/{id}", p.ObterProduto).Methods("GET")
	r.HandleFunc("/produto/{id}", p.RemoverProduto).Methods("DELETE")
	r.HandleFunc("/produtos", p.ObterTodosProdutos).Methods("GET")
	r.HandleFunc("/pedido", p.IncluirPedido).Methods("POST")
	r.HandleFunc("/pedidos", p.ObterPedidosAtivos).Methods("GET")
	r.HandleFunc("/abrir", p.AbrirLoja).Methods("POST")
	r.HandleFunc("/fechar", p.FecharLoja).Methods("POST")
	r.HandleFunc("/metricas", p.ObterMetricas).Methods("GET")
	go p.ProcessarPedidos()
	http.ListenAndServe(":8080", r)
}
