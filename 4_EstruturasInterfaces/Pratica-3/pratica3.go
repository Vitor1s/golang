package main

import (
	"fmt"
	"math" // vai ajudar a calcular a raiz quadrada em resumo area de quadrado e nao aproximar pra 3
)

type geometria interface {
	area() float64
	perimetro() float64
}

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (qua quadrado) area() float64 {
	return qua.lado * qua.lado
}

func (qua quadrado) perimetro() float64 {
	return 4 * qua.lado
}

func (circ circulo) area() float64 {
	return math.Pi * circ.raio * circ.raio
}

func (circ circulo) perimetro() float64 {
	return 2 * math.Pi * circ.raio
}

func medir(g geometria) {
	fmt.Println("Forma: ", g)
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimetro: ", g.perimetro())
}

func main() {
	qua := quadrado{lado: 10}
	circ := circulo{raio: 5}

	medir(qua)
	medir(circ)
}
