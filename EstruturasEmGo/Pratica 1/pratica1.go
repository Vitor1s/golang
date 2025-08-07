package main

import "fmt"

//TODO.  Nessa pratica em estruturas vamos ver instrução For

func main() {
	// crazes()
	forisnt()
	// forisnt2()

	// fmt.Println("1")
	// fmt.Println("2")
	// fmt.Println("3")
	// fmt.Println("4")
	// fmt.Println("5")
	// fmt.Println("6")
	// fmt.Println("7")
	// fmt.Println("8")
	// fmt.Println("9")
	// fmt.Println("10")

}

// //TODO. temos outra forma de realizar essa mesma impressao das informações usando `` crazes

// func crazes() {
// 	fmt.Println(`
// 	1
// 	2
// 	3
// 	4
// 	5
// 	6
// 	7
// 	8
// 	9
// 	10`)
// }

func forisnt() {
	i := 1
	for i <= 15 {
		fmt.Println(i)
		i = i + 1
	}

}

// func forisnt2() {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Println(i)
// 	}

// }

// TODO. O for em Go é como um contador automático que faz algo várias vezes enquanto uma condição for verdadeira. De forma simples ele soma o "i" toda vez que identifica que é verdadeiro 

// func exemplo() {
// 	i := 1                      // começa em 1
// 	for i <= 5 {                // enquanto i for menor ou igual a 5
// 		fmt.Println(i)         // imprime o valor de i
// 		i = i + 1              // soma 1 no i
// 	}
// }