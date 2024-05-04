package IANZINHO

import (
	"IANZINHO/modelos/metricas"
)

type ListaProdutos []Produto

func (lp *ListaProdutos) AdicionarProduto(produto Produto) {
	*lp = append(*lp, produto)
	metricas.MMetricas.TotalProdutos++
}

func (lp *ListaProdutos) RemoverProduto(id int) {
	for i, produto := range *lp {
		if produto.ID == id {
			*lp = append((*lp)[:i], (*lp)[i+1:]...)
			metricas.MMetricas.TotalProdutos--
			break
		}
	}
}

func (lp *ListaProdutos) BuscarProdutoByID(id int) (Produto, bool) {
	for _, produto := range *lp {
		if produto.ID == id {
			return produto, true
		}
	}
	return Produto{}, false
}

func (lp ListaProdutos) ListarProdutos() []Produto {
	return lp
}
