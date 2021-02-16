package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Wesley"
	var idade = 36
	versao := 1.1

	fmt.Println("Bom dia", nome, "Idade:", idade)
	fmt.Println("Versão", versao)

	fmt.Println("variávei nome é", reflect.TypeOf(nome))
	fmt.Println("variávei idade é", reflect.TypeOf(idade))
	fmt.Println("variávei versão é", reflect.TypeOf(versao))
}
