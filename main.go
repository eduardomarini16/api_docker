package main

import (
	"database/sql"
	"fmt"
)

type Produto struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func main() {
	// Conecta ao banco de dados PostgreSQL
	db, err := sql.Open("postgres", "host=localhost user=postgres password=postgres dbname=produtos_db sslmode=disable")
	if err != nil {
		fmt.Println("Erro ao conecta o banco de dados:", err)
	}
	defer db.Close()

	// Cri as tabelas no banco de dados
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS produtos (
		id SERIAL PRIMARY KEY,
		nome VARCHAR(100) NOT NULL
	)`)
	if err != nil {
		fmt.Println("Erro ao criar tabela produtos:", err)
	}
}
