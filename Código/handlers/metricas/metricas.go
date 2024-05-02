package metricas

import (
	"encoding/json"
	m "modelos/metricas/metricas"
	"net/http"
)

var metricas m.Metricas

func handleObterMétricas(w http.ResponseWriter, r *http.Request) {
	atualizarMétricas()
	json.NewEncoder(w).Encode(metricas)
}

func atualizarMétricas() {
	metricas.TotalProdutos = len(listaProdutos)
	metricas.PedidosAndamento = len(filaPedidos)
}
