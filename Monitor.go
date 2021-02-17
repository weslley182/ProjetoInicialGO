package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
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
	const qtdMonitoramento = 5
	const tempoDelay = 3

	fmt.Println("Monitorando...")
	//sites := []string{"https://github.com/weslley182/ProjetoInicialGO",
	//	"https://random-status-code.herokuapp.com"}

	sites := lerSitesDoArquivo()

	for i := 0; i < qtdMonitoramento; i++ {
		testarSites(sites)
		time.Sleep(tempoDelay * time.Second)
	}
	fmt.Println()
}

func testarSites(sites []string) {
	const statusOkIni = 200
	const statusOkFim = 300

	for _, site := range sites {
		resp, erro := http.Get(site)

		if erro != nil {
			fmt.Println("Erro ao verificar site", erro)
		}

		if resp.StatusCode >= statusOkIni && resp.StatusCode < statusOkFim {
			fmt.Println("O site'", site, "'esta em funcionamento", resp.StatusCode)
		} else {
			fmt.Println("O site'", site, "'esta com problemas", resp.StatusCode)
		}
	}
	fmt.Println()
}

func lerSitesDoArquivo() []string {
	const delimitadorQuebraLinha = '\n'

	var sites []string
	arquivo, erro := os.Open("sites.txt")

	if erro != nil {
		fmt.Println("Erro ao verificar aquivo", erro)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, erro := leitor.ReadString(delimitadorQuebraLinha)
		sites = append(sites, strings.TrimSpace(linha))
		if erro == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func iniciarLog() {
	fmt.Println(("Verificando Logs..."))
}
