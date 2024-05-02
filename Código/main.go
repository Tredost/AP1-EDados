package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Produto struct {
	ID        int     `json:"id"`
	Nome      string  `json:"nome"`
	Descrição string  `json:"descrição"`
	Valor     float64 `json:"valor"`
}

type Pedido struct {
	ID         int                 `json:"id"`
	Delivery   bool                `json:"delivery"`
	Produtos   []QuantidadeProduto `json:"produtos"`
	ValorTotal float64             `json:"valor_total"`
}

type QuantidadeProduto struct {
	ID         int `json:"id"`
	Quantidade int `json:"quantidade"`
}

type Métricas struct {
	TotalProdutos     int     `json:"total_produtos"`
	PedidosEncerrados int     `json:"pedidos_encerrados"`
	PedidosAndamento  int     `json:"pedidos_andamento"`
	FaturamentoTotal  float64 `json:"faturamento_total"`
}

type ListaDeProdutos struct {
	Produtos []Produto
}

type FilaDePedidos struct {
	Pedidos []Pedido
}

var (
	ListaProdutos ListaDeProdutos
	FilaPedidos   FilaDePedidos
	produtoID     int = 1
	pedidoID      int = 1
	métricas      Métricas
	lojaAberta    bool
)

// recebe produto e coloca na lsita
func (lp *ListaDeProdutos) AdicionarProduto(produto Produto) {
	lp.Produtos = append(lp.Produtos, produto)
}

// tira da lista pelo id
func (lp *ListaDeProdutos) RemoverProduto(id int) {
	for i, produto := range lp.Produtos {
		if produto.ID == id {
			lp.Produtos = append((lp.Produtos)[:i], (lp.Produtos)[i+1:]...)
			break
		}
	}
}

func (lp ListaDeProdutos) BuscarProdutoByID(id int) (Produto, bool) {
	for _, produto := range lp.Produtos {
		if produto.ID == id {
			return produto, true
		}
	}
	return Produto{}, false
}

func (lp ListaDeProdutos) ListarProdutos() []Produto {
	return lp.Produtos
}

func (fp *FilaDePedidos) IncluirPedido(pedido Pedido) {
	FilaPedidos.Pedidos = append(FilaPedidos.Pedidos, pedido)
}

func (fp *FilaDePedidos) PedidosEmAberto() FilaDePedidos {
	var pedidosAbertos FilaDePedidos
	for _, pedido := range fp.Pedidos {
		if pedido.Delivery || len(pedido.Produtos) > 0 { // mexer depois
			pedidosAbertos.Pedidos = append(pedidosAbertos.Pedidos, pedido)
		}
	}
	return pedidosAbertos
}

func (fp *FilaDePedidos) ExpedirPedido() {
	if len(fp.Pedidos) > 0 {
		fp.Pedidos = fp.Pedidos[1:] // remove o primeiro pedido da fila
		métricas.PedidosEncerrados++
	}
}

func handleAdicionarProduto(w http.ResponseWriter, r *http.Request) {
	var Produto Produto
	err := json.NewDecoder(r.Body).Decode(&Produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Produto.ID = produtoID
	produtoID++
	ListaProdutos.AdicionarProduto(Produto)
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
	produto, encontrado := ListaProdutos.BuscarProdutoByID(id)
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
	ListaProdutos.RemoverProduto(id)
	w.WriteHeader(http.StatusNoContent)
}

func handleObterTodosProdutos(w http.ResponseWriter, r *http.Request) {
	produtos := ListaProdutos.ListarProdutos()
	if len(produtos) == 0 {
		fmt.Fprintln(w, "Não há produtos cadastrados")
		return
	}
	json.NewEncoder(w).Encode(produtos)
}

func handleIncluirPedido(w http.ResponseWriter, r *http.Request) {
	var pedido Pedido
	err := json.NewDecoder(r.Body).Decode(&pedido)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var valorTotal float64
	for _, qp := range pedido.Produtos {
		produto, encontrado := ListaProdutos.BuscarProdutoByID(qp.ID)
		if !encontrado {
			http.Error(w, fmt.Sprintf("Produto com ID %d não encontrado", qp.ID), http.StatusBadRequest)
			return
		}
		valorTotal += produto.Valor * float64(qp.Quantidade)
	}

	pedido.ID = pedidoID
	pedidoID++

	if pedido.Delivery {
		valorTotal += 10
	}
	pedido.ValorTotal = valorTotal

	FilaPedidos.IncluirPedido(pedido)
	métricas.FaturamentoTotal += pedido.ValorTotal

	w.WriteHeader(http.StatusCreated)
}

func handleObterPedidosAtivos(w http.ResponseWriter, r *http.Request) {
	pedidosAtivos := FilaPedidos.PedidosEmAberto()
	if len(pedidosAtivos.Pedidos) == 0 {
		fmt.Fprintln(w, "Não há pedidos ativos")
		return
	}
	json.NewEncoder(w).Encode(pedidosAtivos)
}

func handleAbrirLoja(w http.ResponseWriter, r *http.Request) {
	lojaAberta = true
	go expedirPedidos()
	fmt.Fprintln(w, "Loja aberta")
}

func handleFecharLoja(w http.ResponseWriter, r *http.Request) {
	lojaAberta = false
	fmt.Fprintln(w, "Loja fechada")
}

func handleObterMétricas(w http.ResponseWriter, r *http.Request) {
	atualizarMétricas()
	json.NewEncoder(w).Encode(métricas)
}

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

func expedirPedidos() {
	for lojaAberta {
		pedidosAtivos := FilaPedidos.PedidosEmAberto()
		if len(pedidosAtivos.Pedidos) > 0 {
			time.Sleep(30 * time.Second)
			FilaPedidos.ExpedirPedido()
		}
	}
}

func atualizarMétricas() {
	métricas.TotalProdutos = len(ListaProdutos.Produtos)
	métricas.PedidosAndamento = len(FilaPedidos.Pedidos)
}
