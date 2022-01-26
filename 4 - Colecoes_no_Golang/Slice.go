package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

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
	// Array
	var site [4]string
	site[0] = "https://www.alura.com.br"
	site[1] = "www.google.com"
	site[2] = "www.youtube.com"

	//Slice
	sites := []string{"www.alura.com.br", "www.google.com", "www.youtube.com"}

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
	}

}

func testaSite(site string) {
	resp, _ := http.Get("http://" + site)

	if resp.StatusCode == 200 {
		fmt.Println("ACESSADO: ", site)
	} else {
		fmt.Println("ERRO: ", resp.StatusCode)
	}
}
