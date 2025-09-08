package main

import "fmt"

/*
🎯 ATIVIDADE 7: FUNÇÕES AVANÇADAS

📝 OBJETIVOS:
- Criar funções com diferentes tipos de parâmetros
- Trabalhar com funções que retornam múltiplos valores
- Usar funções variádicas (parâmetros variáveis)
- Implementar funções recursivas
- Praticar reutilização de código

✅ TAREFAS:
*/

func exercicioFuncoes() {
	fmt.Println("=== ATIVIDADE 7: FUNÇÕES AVANÇADAS ===")

	// TODO 1: Funções matemáticas básicas
	// Implemente funções que retornam resultado e erro:
	// - dividir(a, b float64) (float64, error) // erro se b == 0
	// - raizQuadrada(n float64) (float64, error) // erro se n < 0
	// - potencia(base, expoente int) int
	// Teste todas as funções com valores válidos e inválidos

	// TODO 2: Função variádica (parâmetros variáveis)
	// - calcularMedia(numeros ...float64) float64
	// - encontrarMaior(numeros ...int) int
	// - concatenarStrings(separador string, palavras ...string) string
	// Teste com diferentes quantidades de parâmetros

	// TODO 3: Funções que retornam múltiplos valores
	// - analisarNumeros(nums []int) (maior, menor, soma int)
	// - validarEmail(email string) (bool, string) // retorna se válido e mensagem
	// - calcularImposto(salario float64) (bruto, liquido, imposto float64)

	// TODO 4: Função recursiva
	// - fatorial(n int) int
	// - fibonacci(n int) int
	// - contarDigitos(n int) int
	// - somaDigitos(n int) int

	// TODO 5: Funções para slice/array
	// - inverterSlice(nums []int) []int
	// - filtrarPares(nums []int) []int
	// - removerDuplicados(nums []int) []int
	// - ordenarSimples(nums []int) []int // implementação bubble sort simples

	// TODO 6: Funções para strings
	// - contarVogais(texto string) int
	// - isPalindromo(palavra string) bool
	// - capitalize(frase string) string // primeira letra maiúscula
	// - inverterString(s string) string

	// TODO 7: Calculadora avançada com funções
	// Crie um sistema que:
	// - menu() // mostra opções
	// - executarOperacao(op string, a, b float64) (float64, error)
	// - historico(operacoes []string) // mostra histórico
	// Operações: +, -, *, /, %, ^(potência)

	// TODO 8: Sistema de validação
	// - validarIdade(idade int) (bool, string)
	// - validarSenha(senha string) (bool, []string) // retorna lista de problemas
	// - validarCPF(cpf string) bool // validação simples
	// - validarTelefone(tel string) bool

	// TODO 9: Processador de dados
	// - processarVendas(vendas []float64) (total, media, maior, menor float64)
	// - gerarRelatorio(dados map[string]float64) string
	// - aplicarDesconto(preco float64, desconto float64) float64
	// - calcularJurosCompostos(capital, taxa float64, tempo int) float64

	// TODO 10: Sistema de conversões
	// - celsiusParaFahrenheit(c float64) float64
	// - quilometrosParaMilhas(km float64) float64
	// - segundosParaHoras(segundos int) (horas, minutos, segundos int)
	// - bytesParaHumano(bytes int64) string // ex: "1.5 MB"

	// ========================================
	// ESCREVA SEU CÓDIGO AQUI EMBAIXO:
	// ========================================

}

/*
🎓 DICAS IMPORTANTES:

1. FUNÇÕES COM MÚLTIPLOS RETORNOS:
   func exemplo() (int, string, bool) {
       return 42, "ok", true
   }
   a, b, c := exemplo()

2. TRATAMENTO DE ERRO:
   resultado, err := funcaoQueDeErro()
   if err != nil {
       // tratar erro
   }

3. FUNÇÕES VARIÁDICAS:
   func soma(nums ...int) int {
       total := 0
       for _, n := range nums {
           total += n
       }
       return total
   }

4. RECURSÃO:
   func fatorial(n int) int {
       if n <= 1 {
           return 1
       }
       return n * fatorial(n-1)
   }

5. PASSAGEM POR VALOR vs REFERÊNCIA:
   func modificar(slice []int) { }     // slice é referência
   func modificar(valor int) { }       // valor é cópia
   func modificar(ptr *int) { }        // ponteiro para modificar

💡 EXEMPLO DE SAÍDA ESPERADA:
=== ATIVIDADE 7: FUNÇÕES AVANÇADAS ===
10 ÷ 2 = 5.00
10 ÷ 0 = erro: divisão por zero
Raiz de 16 = 4.00
Raiz de -4 = erro: número negativo
Média de 1,2,3,4,5 = 3.00
Maior entre 5,2,8,1 = 8
Concatenação: "Go-é-incrível"
Análise: maior=10, menor=1, soma=55
Email válido: true, "formato correto"
Salário: bruto=5000, líquido=4000, imposto=1000
Fatorial de 5 = 120
Fibonacci(7) = 13
Slice invertido: [5,4,3,2,1]
Números pares: [2,4,6,8]
Vogais em 'programação' = 5
'radar' é palíndromo: true
Idade 25: válida, "idade dentro do limite"
Senha 'abc': inválida, ["muito curta", "sem números"]
25°C = 77°F
100km = 62.14 milhas
3665 segundos = 1h 1m 5s
*/

// Para testar: descomente a linha abaixo
// func main() { exercicioFuncoes() }
