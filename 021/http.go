package main

import (
	"log"
	"net/http"
)

func main() {
	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

	// CLIENTE - SERVIDOR

	// Request - Response

	// Rotas
	// 		URI - Identificador do Recurso
	// 		Método - GET, POST, PUT, DELETE, PATCH, OPTIONS

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Olá Mundo!</h1>"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Carregar página de usuários!</h1>"))
}
