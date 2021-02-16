package main

import (
	"fmt"
	"os"
)

func main() {
	exibeIntroducao()

	exibeMenu()

	comando := leComando()
	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println(("Logs..."))
	case 0:
		fmt.Println("Saindo...")
		os.Exit(0)
	default:
		fmt.Println("Erro...")
		os.Exit(-1)
	}
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
	fmt.Println()
}

func exibeIntroducao() {
	nome := "Wesley"
	versao := 1.1

	fmt.Println("Bom dia", nome)
	fmt.Println("Versão", versao)
	fmt.Println()
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	fmt.Println("Valor da variável: ", comandoLido)
	fmt.Println("Endereço da variável:", &comandoLido)
	fmt.Println()

	return comandoLido
}
