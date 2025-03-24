package main

import (
	"log"
	"net/http"
	"oauth/routecontrollers" // Certifique-se que o caminho do pacote está correto
)

func main() {
	router := routecontrollers.Rotinhadesgracada() // Nome da função deve começar com letra maiúscula

	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
