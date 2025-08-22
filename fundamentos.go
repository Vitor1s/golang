/*
================================================================================
FUNDAMENTOS EM GO
====================================================≠=======================
Este arquivo demonstra os conceitos fundamentais da linguagem Go de forma
estruturada e didática, incluindo:
- Tipos básicos (string, int, float, bool)
- Operações com strings
- Operadores lógicos (booleanos)
- Declaração de variáveis
- Structs (estruturas)
- Boas práticas de formatação
================================================================================
*/

package main // Define que este é o pacote principal (executável)

import "fmt" // Importa o pacote fmt para funções de entrada/saída

// Função principal - ponto de entrada do programa
func main() {
	// ===================================================================
	// 1. TIPOS BÁSICOS E OPERAÇÕES FUNDAMENTAIS
	// ===================================================================
	
	fmt.Println("=== FUNDAMENTOS EM GO ===")
	fmt.Println()
	
	// String básica - tipo texto
	fmt.Println("Hello, World") // Imprime uma string literal
	
	// Operações aritméticas - Go suporta operações matemáticas básicas
	fmt.Println("2+3 =", 2+3) // Soma de dois inteiros
	
	fmt.Println()
	fmt.Println("=== OPERAÇÕES COM STRINGS ===")
	
	// len() - função que retorna o comprimento de uma string
	fmt.Println("Comprimento de 'Hello, World':", len("Hello, World"))
	
	// Acesso a caracteres por índice - retorna o código ASCII/Unicode
	// Importante: índices começam em 0
	fmt.Println("Terceiro caractere de 'Hello World' (índice 2):", "Hello World"[2])
	
	// Concatenação de strings usando o operador +
	fmt.Println("Concatenação:", "Hello "+"World")
	
	// ===================================================================
	// 2. OPERADORES LÓGICOS (BOOLEANOS)
	// ===================================================================
	
	fmt.Println()
	fmt.Println("=== OPERADORES LÓGICOS ===")
	
	// Operador AND (&&) - retorna true apenas se ambos forem true
	fmt.Println("true && true =", true && true)   // true
	fmt.Println("true && false =", true && false) // false
	
	// Operador OR (||) - retorna true se pelo menos um for true
	fmt.Println("true || true =", true || true)   // true
	fmt.Println("true || false =", true || false) // true
	
	// Operador NOT (!) - inverte o valor booleano
	fmt.Println("!true =", !true) // false
	
	// ===================================================================
	// 3. DECLARAÇÃO E USO DE VARIÁVEIS
	// ===================================================================
	
	fmt.Println()
	fmt.Println("=== DECLARAÇÃO DE VARIÁVEIS ===")
	
	// Declaração de variáveis com inferência de tipo
	// Go automaticamente detecta o tipo baseado no valor atribuído
	var nome = "Vitor"     // string
	var idade = 25         // int
	var versao = 1.0       // float64
	
	// Uso das variáveis em fmt.Println
	fmt.Println("Olá", nome, "vi que tem", idade, "anos")
	fmt.Println("Nosso teste está no início, na versão", versao)
	
	// Outra forma de declarar variável
	var mensagem = "Hello World"
	fmt.Println("Mensagem:", mensagem)
	
	// ===================================================================
	// 4. STRUCTS - ESTRUTURAS DE DADOS
	// ===================================================================
	
	fmt.Println()
	fmt.Println("=== ESTRUTURAS (STRUCTS) ===")
	
	// Definição de um tipo personalizado (struct)
	// Struct é uma coleção de campos que permitem agrupar dados relacionados
	type Pessoa struct {
		nome   string  // Campo do tipo string
		idade  int     // Campo do tipo int
		altura float32 // Campo do tipo float32
	}
	
	// Criação de uma instância da struct Pessoa
	// Ordem dos valores deve corresponder à ordem dos campos na definição
	pessoa := Pessoa{"Vitor", 25, 1.80}
	
	// Exibição da struct completa
	fmt.Println("Dados da pessoa:", pessoa)
	
	// Acesso a campos individuais da struct
	fmt.Println("Nome:", pessoa.nome)
	fmt.Println("Idade:", pessoa.idade)
	fmt.Println("Altura:", pessoa.altura, "m")
	
	// ===================================================================
	// 5. DIFERENTES FORMAS DE DECLARAR VARIÁVEIS
	// ===================================================================
	
	fmt.Println()
	fmt.Println("=== FORMAS DE DECLARAÇÃO ===")
	
	// Forma 1: var com tipo explícito
	var cidade string = "São Paulo"
	fmt.Println("Cidade (var explícita):", cidade)
	
	// Forma 2: var com inferência de tipo
	var pais = "Brasil"
	fmt.Println("País (var inferida):", pais)
	
	// Forma 3: declaração curta (apenas dentro de funções)
	regiao := "Sudeste"
	fmt.Println("Região (declaração curta):", regiao)
	
	// ===================================================================
	// 6. CONSTANTES
	// ===================================================================
	
	fmt.Println()
	fmt.Println("=== CONSTANTES ===")
	
	// Constantes são valores que não podem ser alterados após a declaração
	const PI = 3.14159
	const PROGRAMA = "Fundamentos Go"
	
	fmt.Println("Valor de PI:", PI)
	fmt.Println("Nome do programa:", PROGRAMA)
	
	fmt.Println()
	fmt.Println("=== FIM DOS FUNDAMENTOS ===")
}


