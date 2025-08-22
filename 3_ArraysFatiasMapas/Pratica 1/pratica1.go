package main


import "fmt"

func main() {
     praticaarray()
}

//TODO. Aqui vamos ver sobre arrays que no codigo são definidar com [] e o tipo de dado que ele vai armazenar. Onde vamos usar tipos complexos como structs.

func praticaarray() {
	// 1) Cria um array chamado x com 5 posições para números decimais
	var x [5]float64

	// 2) Coloca um número em cada posição do array
	x[0] = 5.3 // primeira posição (índice 0)
	x[1] = 6.3 // segunda posição (índice 1)
	x[2] = 7.3 // terceira posição (índice 2)
	x[3] = 8.3 // quarta posição (índice 3)
	x[4] = 9.3 // quinta posição (índice 4)

	// 3) Mostra todo o array na tela (aparece: [5.3 6.3 7.3 8.3 9.3])
	fmt.Println(x)

	// 4) Cria a variável 'total' para guardar a soma; começa em zero
	var total float64 = 0

	// 5) Percorre o array do começo ao fim somando cada valor em 'total'
	//    len(x) é usado para pegar a quantidade de posições do array (aqui é 5)
	for i := 0; i < len(x); i++ {
		total += x[i] // pega o valor da posição i e adiciona ao total
	}

	// 6) Calcula a média: divide o total pela quantidade de elementos
	media := total / float64(len(x))

	// 7) Mostra a média na tela (com esse exemplo o resultado será 7.3)
	fmt.Println("Média:", media)
}