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

	// Atualizando registros (UPDATE)
	stmt, err := db.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")
	if err != nil {
		log.Fatal("Erro ao preparar a instrução UPDATE:", err)
	}
	defer stmt.Close() // Garante que a declaração seja fechada no final

	// Executando o UPDATE
	if _, err := stmt.Exec("Uóxiton Istive", 1); err != nil {
		log.Fatal("Erro ao executar o UPDATE para o ID 1:", err)
	}
	if _, err := stmt.Exec("Sheristone Uasleska", 2); err != nil {
		log.Fatal("Erro ao executar o UPDATE para o ID 2:", err)
	}

	// Deletando registros (DELETE)
	stmt2, err := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if err != nil {
		log.Fatal("Erro ao preparar a instrução DELETE:", err)
	}
	defer stmt2.Close() // Garante que a declaração seja fechada no final

	// Executando o DELETE
	if _, err := stmt2.Exec(3); err != nil {
		log.Fatal("Erro ao executar o DELETE para o ID 3:", err)
	}

	log.Println("Operações de UPDATE e DELETE realizadas com sucesso!")
}
