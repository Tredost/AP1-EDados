package metricas

type Metricas struct {
	TotalProdutos     int     `json:"total_produtos"`
	PedidosEncerrados int     `json:"pedidos_encerrados"`
	PedidosAndamento  int     `json:"pedidos_andamento"`
	FaturamentoTotal  float64 `json:"faturamento_total"`
}
