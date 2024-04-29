package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"strconv"
)

type Produto struct {
	ID          int     `json:"id"`
	Nome        string  `json:"nome"`
	Descrição   string  `json:"descrição"`
	Valor       float64 `json:"valor"`
}

type Pedido struct {
	ID          int                    `json:"id"`
	Delivery    bool                   `json:"delivery"`
	Produtos    []QuantidadeProduto    `json:"produtos"`  // ver se ta certo 
	ValorTotal  float64                `json:"valor_total"`
}

type QuantidadeProduto struct {
	ID         int `json:"id"`  // puxa produto ver se ta certo
	Quantidade int `json:"quantidade"`
}

type Métricas struct {
	TotalProdutos      int     `json:"total_produtos"`
	PedidosEncerrados  int     `json:"pedidos_encerrados"`
	PedidosAndamento   int     `json:"pedidos_andamento"`
	FaturamentoTotal   float64 `json:"faturamento_total"`
}

type ListaProdutos []Produto  // ver se isso ta certo

type FilaPedidos []Pedido // ver se isso ta certo

var (
	listaProdutos ListaProdutos
	filaPedidos   FilaPedidos
	produtoID     int = 1 // ver se ta certo
	pedidoID      int = 1 // ver se ta certo
	métricas      Métricas
	lojaAberta    bool
)

// recebe produto e coloca na lsita
func (lp *ListaProdutos) AdicionarProduto(produto Produto) {
	*lp = append(*lp, produto)
}

// tira da lista pelo id 
func (lp *ListaProdutos) RemoverProduto(id int) {
	for i, produto := range *lp {
		if produto.ID == id {
			*lp = append((*lp)[:i], (*lp)[i+1:]...)
			break
		}
	}
}

func (lp ListaProdutos) BuscarProdutoByID(id int) (Produto, bool) {
	for _, produto := range lp {
		if produto.ID == id {
			return produto, true
		}
	}
	return Produto{}, false
}

func (lp ListaProdutos) ListarProdutos() ListaProdutos {
	return lp
}

func (fp *FilaPedidos) IncluirPedido(pedido Pedido) {
	*fp = append(*fp, pedido)
}

func (fp FilaPedidos) PedidosEmAberto() FilaPedidos {
	var pedidosAbertos FilaPedidos
	for _, pedido := range fp {
		if pedido.Delivery || len(pedido.Produtos) > 0 {  // mexer depois
			pedidosAbertos = append(pedidosAbertos, pedido)
		}
	}
	return pedidosAbertos
}

func (fp *FilaPedidos) ExpedirPedido() {
	if len(*fp) > 0 {
		*fp = (*fp)[1:] // remove o primeiro pedido da fila
		métricas.PedidosEncerrados++
	}
}

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

func handleIncluirPedido(w http.ResponseWriter, r *http.Request) {
	var pedido Pedido
	err := json.NewDecoder(r.Body).Decode(&pedido)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !lojaAberta { 
		fmt.Fprintln(w, "Loja não está aberta")
		return 
	} // VER COM O VICTOR

	// quantidades do item calcucolo
	var valorTotal float64
	for _, qp := range pedido.Produtos {
		produto, encontrado := listaProdutos.BuscarProdutoByID(qp.ID)
		if !encontrado {
			http.Error(w, fmt.Sprintf("Produto com ID %d não encontrado", qp.ID), http.StatusBadRequest)
			return
		}
		valorTotal += produto.Valor * float64(qp.Quantidade)
	}

	pedido.ID = pedidoID
	pedidoID++

	if pedido.Delivery { valorTotal += 10 }
	pedido.ValorTotal = valorTotal

	filaPedidos.IncluirPedido(pedido)
	métricas.FaturamentoTotal += pedido.ValorTotal

	w.WriteHeader(http.StatusCreated)
}

func handleObterPedidosAtivos(w http.ResponseWriter, r *http.Request) {
	pedidosAtivos := filaPedidos.PedidosEmAberto()
	if len(pedidosAtivos) == 0 {
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
		pedidosAtivos := filaPedidos.PedidosEmAberto()
		if len(pedidosAtivos) > 0 {
			time.Sleep(30 * time.Second)
			filaPedidos.ExpedirPedido()
		}
	}
}

func atualizarMétricas() {
	métricas.TotalProdutos = len(listaProdutos)
	métricas.PedidosAndamento = len(filaPedidos)
}