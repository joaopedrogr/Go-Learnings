package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Abrindo a conexão com o banco de dados
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Preparando a declaração SQL para inserção
	stmt, err := db.Prepare("insert into usuarios(nome) values(?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Executando os inserts
	if _, err := stmt.Exec("Maria"); err != nil {
		panic(err)
	}
	if _, err := stmt.Exec("João"); err != nil {
		panic(err)
	}

	// Executando o insert para "Pedro" e pegando o ID do último registro
	res, err := stmt.Exec("Pedro")
	if err != nil {
		panic(err)
	}

	// Recuperando o ID do último registro inserido
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("ID do último registro inserido:", id)

	// Recuperando o número de linhas afetadas
	linhas, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Número de linhas afetadas:", linhas)
}
