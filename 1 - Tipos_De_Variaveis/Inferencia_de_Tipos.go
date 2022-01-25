package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Eladio"
	idade := 20
	versao := 1.2
	fmt.Println("Hello, sr.", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Este programa está na versão:", versao)

	fmt.Println("O tipo da variavel nome é:", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel versão é:", reflect.TypeOf(versao))
}
