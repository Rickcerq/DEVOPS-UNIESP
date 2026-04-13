package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// A conexão com o banco seria usada aqui
	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println("Conectando ao banco:", dbURL)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Sistema Operante"))
	})

	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Soma é uma função simples para o Teste Unitário
func Soma(a, b int) int {
	return a + b
}
