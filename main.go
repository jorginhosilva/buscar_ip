package main

import (
	"buscar_ip/app"
	"fmt"
	"log"
	"os"
)


func main() {
	fmt.Println("Ponto de Partida")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}