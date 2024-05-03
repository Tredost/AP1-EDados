package pedidos

import (
	"net/http"
)

func HandleIncluirPedido(w http.ResponseWriter, r *http.Request, metricas *Metricas) {
	IncluirPedido(w, r, metricas)
}

func HandleObterPedidosAtivos(w http.ResponseWriter, r *http.Request) {
	ObterPedidosAtivos(w, r)
}
