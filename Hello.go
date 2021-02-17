package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			iniciarLog()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Erro...")
			os.Exit(-1)
		}
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

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://github.com/weslley182/ProjetoInicialGO",
		"https://random-status-code.herokuapp.com"}
	for i := 0; i < len(sites); i++ {
		resp, _ := http.Get(sites[i])

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			fmt.Println("O site'", sites[i], "'esta em funcionamento", resp.StatusCode)
		} else {
			fmt.Println("O site'", sites[i], "'esta com problemas", resp.StatusCode)
		}
	}

	fmt.Println()
}

func iniciarLog() {
	fmt.Println(("Verificando Logs..."))
}
