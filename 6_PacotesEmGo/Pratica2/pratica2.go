package main

import (
	// "fmt"
	// "string"
	"io"
	"log"
	"os"
	"bytes"
	"fmt"
)
//Aqui vemos os pacotes io, log e os, todos focados em manipulação de arquivos e saída de dados. 
//Usei uma das várias funções do pacote io que é a função WriteString, que escreve uma string 
//em qualquer implementação de io.Writer (como console, arquivo, etc.). 
//No meu caso, estou escrevendo no console (os.Stdout). 
//Em um cenário real, você poderia usar para escrever em arquivos, 
//imprimir erros ou mensagens de sucesso em logs.
func main(){
	if _, err := io.WriteString(os.Stdout, "Hello, World!\n"); err != nil {
		log.Fatal(err)
	}
	bytes2()
}

/// aqui veremos o pacote bytes que é um pacote que contém funções para manipulação de bytes.
//Neste exemplo, estou usando a função NewReader para ler o conteúdo de um arquivo.
//Em um cenário real, você poderia usar para ler arquivos, 
//processar dados binários ou implementar protocolos de rede.
func bytes2(){
	arquivo, err := os.Open("meus_logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer arquivo.Close()

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(arquivo)
	fmt.Println(buffer.String())
}
