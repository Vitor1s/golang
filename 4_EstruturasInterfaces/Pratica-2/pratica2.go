package main

import "fmt"

// // Definindo uma struct chamada retangulo
// // Structs agrupam dados relacionados em um único tipo
// // Imagine que você tem um retângulo, como uma folha de papel.
// // Ele tem dois lados: um lado comprido (comprimento) e um lado mais curto (altura).
// type retangulo struct {
//     comprimento, altura int
// }

// // Método associado à struct retangulo
// // Um método é uma função com um "receiver" (neste caso, *retangulo)
// // O receiver indica para qual tipo o método pertence
// // Aqui, usamos ponteiro (*retangulo) para acessar/modificar o valor original
// // Aqui, criamos uma "receita" para descobrir quantos quadradinhos cabem dentro do retângulo.
// // Isso é a área: largura vezes altura.
// func (r *retangulo) area() int {
//     return r.comprimento * r.altura
// }

// // Outro método associado à struct retangulo
// // Aqui usamos valor (retangulo), pois não vamos modificar o struct
// // Aqui, criamos uma "receita" para saber o tamanho da borda do retângulo.
// // Isso é o perímetro: soma todos os lados.
// func (r retangulo) perimetro() int {
//     return 2*r.comprimento + 2*r.altura
// }

// func main() {
//     // Criando uma instância da struct retangulo
//     // Agora, vamos desenhar um retângulo com 10 de comprimento e 5 de altura.
//     r := retangulo{comprimento: 10, altura: 5}

//     // Acessando campos diretamente
//     // Vamos mostrar o tamanho dos lados.
//     fmt.Println("Comprimento:", r.comprimento) // lado maior
//     fmt.Println("Altura:", r.altura)           // lado menor

//     // Chamando métodos: note o uso dos parênteses
//     // Agora, vamos usar nossas receitas para descobrir:
//     fmt.Println("Área:", r.area())         // quantos quadradinhos cabem dentro
//     fmt.Println("Perímetro:", r.perimetro()) // quanto mede a borda toda
// }


type Bicicleta struct{
	modelo string
	aro int
	marchas int
}

func (b Bicicleta) detalhes() string {
	return fmt.Sprintf("Modelo: %s, Aro: %d, Marcha: %d", b.modelo, b.aro, b.marchas)
}

func (b Bicicleta) velocidadeMaxima() int{ 
	return b.aro  b.marchas

}

func main() {
	bicicleta := Bicicleta{modelo: "Caloi", aro: 26, marchas: 18}
	fmt.Println(bicicleta.detalhes())
	fmt.Println("Velocidade Máxima:", bicicleta.velocidadeMaxima(), "km/h")
}

