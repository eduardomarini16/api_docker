package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Produto struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

func getProdutos(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM produtos")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()
		var produtos []Produto
		for rows.Next() {
			var produto Produto
			err := rows.Scan(&produto.ID, &produto.Nome)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			produtos = append(produtos, produto)
		}
		c.JSON(http.StatusOK, produtos)
	}
}

func createProduto(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var produto Produto
		err := c.ShouldBindJSON(&produto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Insira o novo produto no banco de dados
		stmt, err := db.Prepare("INSERT INTO produtos (nome) VALUES ($1)")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer stmt.Close()
		_, err = stmt.Exec(produto.Nome)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "produto criado com sucesso"})
	}
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

	router := gin.Default()
	// rotas API
	router.GET("/produtos", getProdutos(db))
	router.POST("/produtos", createProduto(db))
	router.Run(":8080")
}
