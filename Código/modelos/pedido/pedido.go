package pedido

import "time"

type Pedido struct {
	ID         int                 `json:"id"`
	Delivery   bool                `json:"delivery"`
	Produtos   []QuantidadeProduto `json:"produtos"`
	ValorTotal float64             `json:"valor_total"`
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
