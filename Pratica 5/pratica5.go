package main

import (
	"fmt"
)

// //TODO. Nessa pratica vamos ver as conversões de tipos,

func main() {
	pratica5()
	loja()
}

func pratica5() {
	var a int = 10
	var b string = "Vitor"

// 	var c float32 = float32(a) // Convertendo int para float32
// 	//TODO. Aqui estamos convertendo o tipo int para float32, isso é possivel
	// porque float32 é um tipo numérico, assim como int. A conversão é
	// feita usando a sintaxe `float32(a)`, onde `a` é a variável do tipo int.
	fmt.Printf("valor de C é:%g Valor de B:%s\n", a, b)

}
// 	//TODO. Exemplo abaixo de usuabilidade para o caso de Conversão de Tipos 

func loja() {
    preco_inteiro := 25      // R$ 25
    preco_com_desconto := float32(preco_inteiro) * 0.9  // 10% desconto
    
    fmt.Printf("Preço original: R$ %d\n", preco_inteiro)      // %d para int
    fmt.Printf("Com desconto: R$ %.2f\n", preco_com_desconto) // %.2f para 2 casas decimais
}

//  Resumo dos "símbolos mágicos":
// %d = números inteiros (1, 2, 100)
// %g ou %f = números decimais (1.5, 3.14)
// %s = texto ("João", "Hello")