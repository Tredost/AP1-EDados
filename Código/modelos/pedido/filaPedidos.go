package IANZINHO

import (
	p "IANZINHO/processamento"
	"time"
)

type FilaPedidos []Pedido

func (fp *FilaPedidos) IncluirPedido(pedido Pedido) {
	*fp = append(*fp, pedido)
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
		*fp = (*fp)[1:] // remove o primeiro pedido da fila
		p.Metricas.PedidosEncerrados++
	}
}

func ExpedirPedidos() {
	for p.LojaAberta {
		pedidosAtivos := FilaPedidos.PedidosEmAberto()
		if len(pedidosAtivos) > 0 {
			time.Sleep(30 * time.Second)
			FilaPedidos.ExpedirPedido()
		}
	}
}
