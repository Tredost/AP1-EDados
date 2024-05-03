package pedido

type Pedido struct {
	ID         int                 `json:"id"`
	Delivery   bool                `json:"delivery"`
	Produtos   []QuantidadeProduto `json:"produtos"`
	ValorTotal float64             `json:"valor_total"`
}

type QuantidadeProduto struct {
	ID         int `json:"id"`
	Quantidade int `json:"quantidade"`
}
