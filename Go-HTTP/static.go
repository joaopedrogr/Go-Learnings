package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler para servir o conteúdo HTML
func htmlHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	    <meta charset="UTF-8">
	    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	    <meta http-equiv="X-UA-Compatible" content="ie=edge">
	    <title>Página usando Go</title>
	</head>
	<body>
	    <h1>Go!!!!</h1>
	</body>
	</html>`

	// Define o cabeçalho de resposta para o tipo de conteúdo como HTML
	w.Header().Set("Content-Type", "text/html")
	// Envia o HTML como resposta
	fmt.Fprint(w, htmlContent)
}

func main() {
	// Define a rota para o handler HTML
	http.HandleFunc("/", htmlHandler)

	// Inicia o servidor
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
