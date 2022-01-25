package main

import "fmt"

func main() {
	nome := "Eladio"
	var versao float32 = 1.2
	fmt.Println("Hello, sr.", nome)
	fmt.Println("Este programa está na versão:", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")

	var comando int

	fmt.Print("Digite sua opção: ")
	//fmt.Scanf("%d", &comando) Explicita o modificador
	fmt.Scan(&comando) // Infere o tipo desejado

	fmt.Println("O comando escolhido foi:", comando)
	fmt.Println("Endereço de comando é:", &comando)
}
