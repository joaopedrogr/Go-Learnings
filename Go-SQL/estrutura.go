package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Função para executar comandos SQL no banco de dados
func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	// Conexão com o banco de dados MySQL
	db, err := sql.Open("mysql", "root:123456@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Executando comandos SQL
	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)
}
