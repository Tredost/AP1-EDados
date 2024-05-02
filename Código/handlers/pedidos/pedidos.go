package pedidos

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleIncluirPedido(w http.ResponseWriter, r *http.Request) {
	var pedido Pedido
	err := json.NewDecoder(r.Body).Decode(&pedido)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

	if pedido.Delivery {
		valorTotal += 10
	}
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
