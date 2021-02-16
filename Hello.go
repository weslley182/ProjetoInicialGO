package main

import "fmt"

func main() {
	nome := "Wesley"
	versao := 1.1

	fmt.Println("Bom dia", nome)
	fmt.Println("Versão", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("Valor da variável: ", comando)
	fmt.Println("Endereço da variável:", &comando)
}
