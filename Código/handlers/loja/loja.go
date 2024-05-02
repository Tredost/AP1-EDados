package loja

import (
	"fmt"
	"net/http"
	//m "IANZINHO/Código/modelos/pedido/pedido"  PQ NÃO VAAAAAAAAAAAAAAAAAAIIII
)

var lojaAberta bool

func handleAbrirLoja(w http.ResponseWriter, r *http.Request) {
	lojaAberta = true
	go m.expedirPedidos()
	fmt.Fprintln(w, "Loja aberta")
}

func handleFecharLoja(w http.ResponseWriter, r *http.Request) {
	lojaAberta = false
	fmt.Fprintln(w, "Loja fechada")
}
