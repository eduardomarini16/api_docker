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
}
