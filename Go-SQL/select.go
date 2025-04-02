package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	// Abrindo a conexão com o banco de dados
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal("Erro ao abrir a conexão com o banco de dados:", err)
	}
	defer db.Close() // Garantir que a conexão seja fechada ao final

	// Consultando usuários com id maior que 3
	rows, err := db.Query("SELECT id, nome FROM usuarios WHERE id > ?", 3)
	if err != nil {
		log.Fatal("Erro ao executar a consulta:", err)
	}
	defer rows.Close() // Garantir que os resultados da consulta sejam fechados ao final

	// Iterando sobre os resultados
	for rows.Next() {
		var u usuario
		// Lendo os valores das colunas e atribuindo às variáveis
		if err := rows.Scan(&u.id, &u.nome); err != nil {
			log.Fatal("Erro ao ler os dados da linha:", err)
		}
		// Imprimindo os dados do usuário
		fmt.Println(u)
	}

	// Verificando se ocorreu algum erro durante a iteração
	if err := rows.Err(); err != nil {
		log.Fatal("Erro durante a iteração dos resultados:", err)
	}
}
