package produto

type ListaProdutos []Produto // ver de fzr um sruct

// recebe produto e coloca na lsita
func (lp *ListaProdutos) AdicionarProduto(produto Produto) {
	*lp = append(*lp, produto)
}

// tira da lista pelo id
func (lp *ListaProdutos) RemoverProduto(id int) {
	for i, produto := range *lp {
		if produto.ID == id {
			*lp = append((*lp)[:i], (*lp)[i+1:]...)
			break
		}
	}
}

func (lp ListaProdutos) BuscarProdutoByID(id int) (Produto, bool) {
	for _, produto := range lp {
		if produto.ID == id {
			return produto, true
		}
	}
	return Produto{}, false
}

func (lp ListaProdutos) ListarProdutos() ListaProdutos {
	return lp
}
