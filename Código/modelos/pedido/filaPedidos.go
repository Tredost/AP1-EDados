package pedido

import "time"

type FilaPedidos struct {
	Pedidos []Pedido
}

func (fp *FilaPedidos) IncluirPedido(pedido Pedido) {
	fp.Pedidos = append(fp.Pedidos, pedido)
}

func (fp *FilaPedidos) ExpedirPedido() {
	for lojaAberta {
		pedidosAtivos := filaPedidos.PedidosEmAberto()
		if len(pedidosAtivos) > 0 {
			time.Sleep(30 * time.Second)
			filaPedidos.ExpedirPedido()
		}
	}
}

func (lp FilaPedidos) PedidosEmAberto() []Pedido {
	var pedidosAbertos []Pedido
	for _, pedido := range lp.Pedidos {
		if pedido.Delivery || len(pedido.Produtos) > 0 {
			pedidosAbertos = append(pedidosAbertos, pedido)
		}
	}
	return pedidosAbertos
}
