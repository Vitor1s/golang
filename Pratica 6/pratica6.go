package main

import "fmt"

func main() {
	pratica6()
}

//TODO. Aqui vamos aprender como descobrir o tipo de uma variavel, as vezes ao ler um codigo esquecemos ou precisamos identificar para o que aquela variavel ta sendo usada, e pra isso usamos o %T no printf, que nos mostra o tipo da variavel, por exemplo:

func pratica6() {

	aguaquente := 212.0 // Definindo a temperatura da água quente em Fahrenheit
	F := aguaquente
	C := (F - 32) * 5 / 9
	fmt.Printf("A temperatura em Celsius é = %v (%T) A temperatura em Fahrenheit é = %v (%T).\n", C, C, F, F)

}

// %v = valor (funciona com qualquer coisa)
// %d = valor (só números inteiros)
// %f = valor (só números decimais)
// %s = valor (só texto)
// %T = tipo (mostra se é int, string, float, etc.)