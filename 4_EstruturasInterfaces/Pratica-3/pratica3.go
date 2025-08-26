package main

import (
	"fmt"
	"math" // vai ajudar a calcular a raiz quadrada em resumo area de quadrado e nao aproximar pra 3
)

// Beleza aqui entendi que criamos a geometria que vai ser a base de calculo de area e perimetro para o quadrado e circulo
type geometria interface {
	area() float64
	perimetro() float64
}

// Aqui e a estrutura do quadrado onde o lado que podemos chamar de variavel(lado pois quadrado tem 4 lados)
type quadrado struct {
	lado float64
}

// Aqui e a estrutura do circulo onde o raio que podemos chamar de variavel (raio pois circulo tem 1 raio)
type circulo struct {
	raio float64
}

// Aqui é onde criamos a funçao pra somar a area do quadrado e na matematica para char a area é lado *(vezes) lado
// Detalhe o qua que é a variavel que colcoamos la do dentro da func principal main, que demos o valor de 10
func (qua quadrado) area() float64 {
	return qua.lado * qua.lado
}

// Aqui e a funçao pra somar o perimetro do quadrado e na matematica para char o perimetro é 4 *(vezes) lado
// Ai voce se pegunta de onde vem esse valor de 4? E o simples fato de na matematica um quadrado tem 4 lados e por isso multiplicamos por 4 e pq 4 * qua.lado? pq pra achar o perimetro temos que somar com o valor que demos para os lados
func (qua quadrado) perimetro() float64 {
	return 4 * qua.lado
}

// Aqui e a funçao pra somar a area do circulo e na matematica para char a area é pi *(vezes) raio *(vezes) raio
// Detalhe o circ que é a variavel que colcoamos la do dentro da func principal main, que demos o valor de 5
// E o math.Pi é uma constante que representa o numero pi e é uma constante em go que chamamos la de cima da biblioteca math
func (circ circulo) area() float64 {
	return math.Pi * circ.raio * circ.raio
}

// Aqui e a funçao pra somar o perimetro do circulo e na matematica para char o perimetro é 2 *(vezes) pi *(vezes) raio
// Detalhe o circ que é a variavel que colcoamos la do dentro da func principal main, que demos o valor de 5
// E o math.Pi é uma constante que representa o numero pi e é uma constante em go que chamamos la de cima da biblioteca math
func (circ circulo) perimetro() float64 {
	return 2 * math.Pi * circ.raio
}

// Aqui é a funcçao que mede a geometria pra mostrar a area e o perimetro do quadrado e do circulo
// Ai voce se pergunta de onde vem o g? Em termos tecnicos pegamos o typo geometria e chamamos de g
func medir(g geometria) {
	fmt.Println("Forma: ", g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimetro: ", g.perimetro())
}

// Aqui é a funcçao principal main que vai rodar o programa pq em Go o programa começa pela func main
// E aquele medir voando ali? E A funçao logo acima  que estamos usando pra mostrar a area e o perimetro do quadrado e do circulo
func main() {
	qua := quadrado{lado: 10}
	circ := circulo{raio: 5}

	medir(qua)
	medir(circ)
}
