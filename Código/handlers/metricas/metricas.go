package IANZINHO

import (
	p "IANZINHO/processamento"
	"net/http"
)

func HandleObterMetricas(w http.ResponseWriter, r *http.Request) {
	p.ObterMetricas(w, r)
}
