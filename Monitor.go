package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
			imprimirLogs()
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
			registraLog(site, true)
		} else {
			fmt.Println("O site'", site, "'esta com problemas", resp.StatusCode)
			registraLog(site, false)
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

func registraLog(site string, status bool) {
	//https://golang.org/src/time/format.go
	const formatoData = "02/01/2006 15:04:05"
	const quebraLinha = "\n"

	arquivo := retornarArquivoDeLog()

	arquivo.WriteString(site + " _online:" + strconv.FormatBool(status) +
		" " + time.Now().Format(formatoData) + quebraLinha)
	arquivo.Close()
}

func retornarArquivoDeLog() *os.File {
	const permissao = 0666

	arquivo, erro := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, permissao)

	if erro != nil {
		fmt.Println("Erro encontrado", erro)
	}
	return arquivo
}

func imprimirLogs() {
	arquivo, erro := ioutil.ReadFile("log.txt")

	if erro != nil {
		fmt.Println("Erro: ", erro)
	}
	fmt.Println(string(arquivo))
}
