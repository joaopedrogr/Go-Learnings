package main

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Abrindo a conexão com o banco de dados
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal("Erro ao abrir conexão com o banco de dados:", err)
	}
	defer db.Close()

	// Iniciando uma transação
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Erro ao iniciar a transação:", err)
	}
	defer tx.Rollback() // Garante que, se algo der errado, a transação seja revertida

	// Preparando a instrução SQL
	stmt, err := tx.Prepare("insert into usuarios(id, nome) values(?, ?)")
	if err != nil {
		log.Fatal("Erro ao preparar a instrução SQL:", err)
	}
	defer stmt.Close() // Garante que a declaração seja fechada no final

	// Executando os inserts
	if _, err := stmt.Exec(4000, "Bia"); err != nil {
		log.Fatal("Erro ao inserir Bia:", err)
	}
	if _, err := stmt.Exec(4001, "Carlos"); err != nil {
		log.Fatal("Erro ao inserir Carlos:", err)
	}

	// Tentando inserir um ID duplicado (deve causar erro)
	_, err = stmt.Exec(1, "Tiago") // chave duplicada
	if err != nil {
		// Em caso de erro, fazemos o rollback da transação
		tx.Rollback()
		log.Fatal("Erro ao inserir Tiago (provavelmente chave duplicada):", err)
	}

	// Se tudo deu certo, confirmamos a transação
	err = tx.Commit()
	if err != nil {
		log.Fatal("Erro ao confirmar a transação:", err)
	}

	log.Println("Transação concluída com sucesso!")
}
