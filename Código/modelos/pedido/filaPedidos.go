package IANZINHO

import (
	"IANZINHO/modelos/metricas"
)

type FilaPedidos []Pedido

var FPedidos FilaPedidos

func (fp *FilaPedidos) IncluirPedido(pedido Pedido) {
	*fp = append(*fp, pedido)
	metricas.MMetricas.PedidosAndamento++
}

func (fp FilaPedidos) PedidosEmAberto() []Pedido {
	var pedidosAbertos []Pedido
	for _, pedido := range fp {
		if pedido.Delivery || len(pedido.Produtos) > 0 {
			pedidosAbertos = append(pedidosAbertos, pedido)
		}
	}
	return pedidosAbertos
}

func (fp *FilaPedidos) ExpedirPedido() {
	if len(*fp) > 0 {
		*fp = (*fp)[1:]
		metricas.MMetricas.PedidosEncerrados++
		metricas.MMetricas.PedidosAndamento--
	}
}
