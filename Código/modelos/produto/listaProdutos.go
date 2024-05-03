package produto

type ListaProdutos struct {
	Produtos []Produto
}

func (lp *ListaProdutos) AdicionarProduto(produto Produto) {
	lp.Produtos = append(lp.Produtos, produto)
}

func (lp *ListaProdutos) RemoverProduto(id int) {
	for i, produto := range lp.Produtos {
		if produto.ID == id {
			lp.Produtos = append((lp.Produtos)[:i], (lp.Produtos)[i+1:]...)
			break
		}
	}
}

func (lp ListaProdutos) BuscarProdutoByID(id int) (Produto, bool) {
	for _, produto := range lp.Produtos {
		if produto.ID == id {
			return produto, true
		}
	}
	return Produto{}, false
}

func (lp ListaProdutos) ListarProdutos() []Produto {
	return lp.Produtos
}
