package IANZINHO

import (
	"IANZINHO/modelos/metricas"
	pe "IANZINHO/modelos/pedido"
	pr "IANZINHO/modelos/produto"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	ListaProdutos pr.ListaProdutos
	FilaPedidos   pe.FilaPedidos
	ProdutoID     int = 1
	PedidoID      int = 1
	LojaAberta    bool
)

//loja

func AbrirLoja(w http.ResponseWriter, r *http.Request) {
	LojaAberta = true
	fmt.Fprintln(w, "Loja aberta")
}

func FecharLoja(w http.ResponseWriter, r *http.Request) {
	LojaAberta = false
	fmt.Fprintln(w, "Loja fechada")
}

// metricas

func ObterMetricas(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metricas.MMetricas)
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
	fmt.Fprintln(w, "Pedido incluído")
	pedido.ID = PedidoID
	PedidoID++

	if pedido.Delivery {
		valorTotal += 10
	}
	pedido.ValorTotal = valorTotal

	pe.FPedidos.IncluirPedido(pedido)
	metricas.MMetricas.FaturamentoTotal += pedido.ValorTotal

	w.WriteHeader(http.StatusCreated)
}

func ObterPedidosAtivos(w http.ResponseWriter, r *http.Request) {
	pedidosAtivos := pe.FPedidos.PedidosEmAberto()
	if len(pedidosAtivos) == 0 {
		fmt.Fprintln(w, "Não há pedidos ativos")
		return
	}
	json.NewEncoder(w).Encode(pedidosAtivos)
}

func ProcessarPedidos() {
	for {
		if LojaAberta {
			pedidosAtivos := pe.FPedidos.PedidosEmAberto()
			if len(pedidosAtivos) > 0 {
				pe.FPedidos.ExpedirPedido()
			}
		}
	}
}

func AdicionarProduto(w http.ResponseWriter, r *http.Request) {
	var Produto pr.Produto
	err := json.NewDecoder(r.Body).Decode(&Produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "Produto adicionado")
	Produto.ID = ProdutoID
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
	produto, encontrado := ListaProdutos.BuscarProdutoByID(id)
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
	fmt.Fprintln(w, "Produto removido")
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
