package pedido

type FilaPedidos []Pedido // ver de fzr um sruct

func (fp *FilaPedidos) IncluirPedido(pedido Pedido) {
	*fp = append(*fp, pedido)
}

func (fp FilaPedidos) PedidosEmAberto() FilaPedidos {
	var pedidosAbertos FilaPedidos
	for _, pedido := range fp {
		if pedido.Delivery || len(pedido.Produtos) > 0 { // mexer depois
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
