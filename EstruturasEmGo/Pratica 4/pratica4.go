package main

import "fmt"

//TODO. Vamos ver Operadores Lógicos Condicionais com If

func main() {
	praticaou()
	praticae() 
}

func praticaou() {
	x := 4
	if x == 2 || x == 3 {
		fmt.Println("x é 2 ou 3")
	}else {
		fmt.Println("x não é 2 nem 3")
	}

}

//TODO. Acima vimos como seria se usassemos a regra de (OR) (||) que pelo menos uma condição precisa ser verdadeira para que a execução do if ocorra

func praticae() {
	x := 6
	if x%2 == 0 && x%3 == 0 {
		fmt.Println("x é multiplo de 2 e 3")
	}else {
		fmt.Println("x não é multiplo de 2 e 3")
	}

}

//TODO. Acima vimos como seria se usassemos a regra de E (AND)(&&) que todas as condições precisam ser verdadeiras para que a execução do if ocorra