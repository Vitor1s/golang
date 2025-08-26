package main

import "fmt"

// Aqui vamos ver o que é ponteiro (ptr) e como ele funciona.  

// func inicial(x *int){
// 	*x = 2000
// } 


// func main(){
// 	x := 10
// 	fmt.Println(x)
// 	inicial(&x)
// 	fmt.Println(x)
// }

/// Aqui a gente vê que o valor de x nao foi alterado, pq a funçao inicial nao retorna nada, e o que é ponteiro? É uma variavel que armazena o endereço de memoria de outra variavel.



// codigo que fiz para ver o resultado da atividade final 
func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main(){
	res := plus(1,2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1,2,3)
	fmt.Println("1+2+3 =", res)
}