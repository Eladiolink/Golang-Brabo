package main

import "fmt"

func main() {
	nome, idade := nomeIdade()
	// nome, _ := nomeIdade() Ignora o parametro retornado no lugar do "_"
	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
}

func nomeIdade() (string, int) {
	return "Eladio", 20
}
