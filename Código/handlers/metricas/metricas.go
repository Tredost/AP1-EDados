package metricas

import (
	"net/http"
	//IMPORTAR PROCESSAMENTO
)

func HandleObterMetricas(w http.ResponseWriter, r *http.Request, metricas *Metricas) {
	ObterMetricas(w, r, metricas)
}
