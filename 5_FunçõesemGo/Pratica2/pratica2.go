package main

import "fmt"

func main() {
	fmt.Println("=== EXEMPLO DE RECURSÃO ===")
	exemploRecursao()

	fmt.Println("\n=== EXEMPLO DE CLOSURES ===")
	exemploClosures()
}

// ==========================================
// RECURSÃO: Função que chama ela mesma
// ==========================================
func exemploRecursao() {
	fmt.Println("Calculando fatorial de 9...")
	fmt.Println("O que é recursão? É uma função que chama ela mesma!")
	fmt.Println("No caso do fatorial: 9! = 9 × 8 × 7 × 6 × 5 × 4 × 3 × 2 × 1")

	resultado := fatorial(9)
	fmt.Printf("Fatorial de 9 = %d\n", resultado)
	fmt.Println("Como funciona? A função fatorial chama ela mesma até chegar no caso base (n == 0)")
}

// Função recursiva que calcula o fatorial
// Ela chama ela mesma até chegar no caso base
func fatorial(n int) int {
	// CASO BASE: quando n == 0, retorna 1 (condição de parada)
	if n == 0 {
		return 1
	}
	// CASO RECURSIVO: n! = n × (n-1)!
	// A função chama ela mesma com n-1
	return n * fatorial(n-1)
}

// ==========================================
// CLOSURES: Função que retorna uma função
// ==========================================
func exemploClosures() {
	fmt.Println("O que é Closure? É uma função que 'captura' variáveis do escopo externo!")
	fmt.Println("A função interna tem acesso às variáveis da função externa")

	contador := 0 // Variável que será acessada pela função interna

	// Criando uma função que "captura" a variável contador
	// Esta é a MAGIA do Closure: a função interna lembra do valor da variável externa
	incrementar := func() {
		contador++ // Acessa e modifica a variável 'contador' da função externa
		fmt.Printf("Contador: %d\n", contador)
	}

	fmt.Println("Chamando a função incrementar várias vezes:")
	fmt.Println("Cada chamada incrementa o mesmo contador (não cria um novo):")
	incrementar() // Contador: 1 - Primeira chamada
	incrementar() // Contador: 2 - Segunda chamada, lembra do valor anterior
	incrementar() // Contador: 3 - Terceira chamada, continua incrementando
}
