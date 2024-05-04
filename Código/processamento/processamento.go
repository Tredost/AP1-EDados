package IANZINHO

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	m "IANZINHO/modelos/metricas"
	pe "IANZINHO/modelos/pedido"
	pr "IANZINHO/modelos/produto"

	"github.com/gorilla/mux"
)

var (
	ListaProdutos pr.ListaProdutos
	FilaPedidos   pe.FilaPedidos
	ProdutoID     int = 1
	PedidoID      int = 1
	Metricas      m.Metricas
	LojaAberta    bool
)

//loja

func AbrirLoja(w http.ResponseWriter, r *http.Request) {
	LojaAberta = true
	go pe.ExpedirPedidos()
	fmt.Fprintln(w, "Loja aberta")
}

func FecharLoja(w http.ResponseWriter, r *http.Request) {
	LojaAberta = false
	fmt.Fprintln(w, "Loja fechada")
}

// metricas

func ObterMetricas(w http.ResponseWriter, r *http.Request) {
	AtualizarMetricas()
	json.NewEncoder(w).Encode(Metricas)
}

func AtualizarMetricas() {
	Metricas.TotalProdutos = len(ListaProdutos)
	Metricas.PedidosAndamento = len(FilaPedidos)
}

//pedidos

func IncluirPedido(w http.ResponseWriter, r *http.Request) {
	var pedido pe.Pedido
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

	pedido.ID = PedidoID
	PedidoID++

	if pedido.Delivery {
		valorTotal += 10
	}
	pedido.ValorTotal = valorTotal

	FilaPedidos.IncluirPedido(pedido) //ERRADOOOOOOOOOOOOOOOOO
	Metricas.FaturamentoTotal += pedido.ValorTotal

	w.WriteHeader(http.StatusCreated)
}

func ObterPedidosAtivos(w http.ResponseWriter, r *http.Request) {
	pedidosAtivos := FilaPedidos.PedidosEmAberto()
	if len(pedidosAtivos) == 0 {
		fmt.Fprintln(w, "Não há pedidos ativos")
		return
	}
	json.NewEncoder(w).Encode(pedidosAtivos)
}

/*func (fp *pe.FilaPedidos) ExpedirPedido() {
	if len(*fp) > 0 {
		*fp = (*fp)[1:] // remove o primeiro pedido da fila
		Metricas.PedidosEncerrados++
	}
}

func ExpedirPedidos() {
	for LojaAberta {
		pedidosAtivos := FilaPedidos.PedidosEmAberto()
		if len(pedidosAtivos) > 0 {
			time.Sleep(30 * time.Second)
			FilaPedidos.ExpedirPedido()
		}
	}
}*/

// produtos

func AdicionarProduto(w http.ResponseWriter, r *http.Request) {
	var Produto pr.Produto
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

func ObterProduto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	produto, encontrado := ListaProdutos.BuscarProdutoByID(id) //TA ERRADO
	if !encontrado {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(produto)
}

func RemoverProduto(w http.ResponseWriter, r *http.Request) {
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

func ObterTodosProdutos(w http.ResponseWriter, r *http.Request) {
	produtos := ListaProdutos.ListarProdutos()
	if len(produtos) == 0 {
		fmt.Fprintln(w, "Não há produtos cadastrados")
		return
	}
	json.NewEncoder(w).Encode(produtos)
}
