package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	for {
		menu()
	}

}

func exibeIntroducao() {
	nome := "Eladio"
	var versao float32 = 1.2
	fmt.Println("Hello, sr.", nome)
	fmt.Println("Este programa está na versão:", versao)
}

func menu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")

	comando := lerComando()

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")

	case 0:
		fmt.Println("Saindo :)")
		os.Exit(0)
	default:
		fmt.Println("Opção não existe :/")
		os.Exit(-1)
	}
}

func lerComando() int {
	var comando int

	fmt.Print("Digite sua opção: ")
	fmt.Scan(&comando)

	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	//fmt.Println(resp.StatusCode)

	if resp.StatusCode == 200 {
		fmt.Println("ACESSADO: ", resp.StatusCode)
	} else {
		fmt.Println("ERRO: ", resp.StatusCode)
	}
}
