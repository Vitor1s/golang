package main

import "fmt"

//TODO. Vamos ver como funciona o "if"(SE) em go

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {   
			fmt.Println(i, "é par")
		} else {
			fmt.Println(i, "é ímpar")
		}
	}
}


//TODO. Como saber se um número é par ou ímpar multiplicando por 2 se o resultado for 0 é par, se não é ímpar

//TODO. A porcentagem seria a divisão do número por 2, se o resto for 0 é par, se não é ímpar