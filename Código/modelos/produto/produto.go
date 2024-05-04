package IANZINHO

type Produto struct {
	ID        int     `json:"id"`
	Nome      string  `json:"nome"`
	Descrição string  `json:"descrição"`
	Valor     float64 `json:"valor"`
}
