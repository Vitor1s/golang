package main

import "fmt"

// Aqui vou vamos entender o que é uma funçao em Go e como ela funciona e como pode ser melhorada ainda mais, voce sabe o que é uma funçao? uma funçao é um bloco de codigo que pode ser reutilizado, ou seja, uma funçao é um bloco de codigo que pode ser reutilizado em outro lugar

//obs: sabe pra que serve o range? o range é uma funçao que retorna o indice e o valor de cada elemento de um slice ou array, resumindo de forma mais ele pega todos os valores de um slice ou array e soma eles

func media(notas []float64) float64{  // aqui criamos é a media que vamos usar para calcular a media das notas
	total := 0.0  // aqui criamos a variavel total que vai armazenar a soma das notas epq zero? pq zero é o valor inicial da variavel total 
	for _, nota := range notas { // aqui criamos o for que vai percorrer o slice notas e somar as notas
		total += nota // aqui somamos as notas
	}
	return total / float64(len(notas)) // Aqui a gente pausa a funçao e fala pra ela mano retorna o total dividido pelo numero de notas pra consegui a media do total
	//obs: sabe o que é o len? o len é uma funçao que retorna o numero de elementos de um slice ou array, resumindo de forma mais ele pega o numero de elementos de um slice ou array e divide pelo numero de notas pra consegui a media do total
}

func main(){
   // aqui criamos o slice notas que vai armazenar as notas do aluno
	notas := []float64{7, 8, 9, 10}
	 fmt.Println(media(notas)) // aqui chamamos a funçao media e passamos o slice notas como parametro


	
}