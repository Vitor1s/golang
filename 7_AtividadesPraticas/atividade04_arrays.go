package main

import "fmt"

/*
🎯 ATIVIDADE 4: ARRAYS E SLICES

📝 OBJETIVOS:
- Trabalhar com arrays de tamanho fixo
- Manipular slices (fatias)
- Praticar operações de slice (append, corte, etc.)
- Calcular estatísticas (média, maior, menor)

✅ TAREFAS:
*/

func exercicioArrays() {
	fmt.Println("=== ATIVIDADE 4: ARRAYS E SLICES ===")

	// TODO 1: Array básico de temperaturas
	// - crie um array [7]float64 para temperatura da semana
	// - preencha com: 23.5, 25.0, 22.8, 24.2, 26.1, 21.9, 23.7
	// - calcule e imprima a média das temperaturas
	// - encontre a maior e menor temperatura

	// TODO 2: Slice de nomes
	// - crie um slice: nomes := []string{"Ana", "Bruno", "Carlos"}
	// - adicione mais 2 nomes usando append
	// - imprima todos os nomes numerados (1. Ana, 2. Bruno, etc.)
	// - imprima quantos nomes tem no total

	// TODO 3: Manipulação de slice numérico
	// - crie: numeros := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	// - extraia os 3 primeiros números (slice[0:3])
	// - extraia os 3 últimos números
	// - extraia os números do meio (posições 3 a 6)
	// - imprima cada sub-slice

	// TODO 4: Análise de vendas
	// - vendas := []float64{1500.50, 2300.75, 1800.00, 2100.25, 1950.80}
	// - calcule o total de vendas
	// - calcule a média de vendas
	// - encontre a maior e menor venda
	// - conte quantas vendas foram acima da média

	// TODO 5: Filtro de números
	// - dados := []int{15, 8, 23, 4, 16, 42, 7, 11, 19, 33}
	// - crie um slice só com números pares
	// - crie um slice só com números maiores que 15
	// - use append para construir os novos slices

	// TODO 6: Simulador de fila
	// - crie um slice vazio: fila := []string{}
	// - simule pessoas chegando: adicione "João", "Maria", "Pedro"
	// - simule atendimento: remova o primeiro da fila (João)
	// - adicione mais pessoas: "Ana", "Carlos"
	// - imprima a fila atual
	// Dica: para remover primeiro elemento: slice = slice[1:]

	// TODO 7: Matriz simples (array de arrays)
	// - crie uma "matriz" 3x3 usando: var matriz [3][3]int
	// - preencha com números sequenciais (1,2,3 na primeira linha, 4,5,6 na segunda, etc.)
	// - imprima a matriz formatada
	// - calcule a soma de todos os elementos

	// TODO 8: Estatísticas avançadas
	// - notas := []float64{7.5, 8.2, 6.8, 9.1, 7.0, 8.5, 6.9, 7.8, 8.0, 7.2}
	// - calcule: média, mediana (valor central), maior, menor
	// - conte quantas notas estão acima/abaixo da média
	// Dica: para mediana, você precisará ordenar (pode fazer manualmente com loops)

	// ========================================
	// ESCREVA SEU CÓDIGO AQUI EMBAIXO:
	// ========================================

}

/*
🎓 DICAS IMPORTANTES:

1. ARRAYS vs SLICES:
   [tamanho]tipo: array (tamanho fixo)
   []tipo: slice (tamanho variável)

2. OPERAÇÕES COM SLICES:
   append(slice, elemento)     // adiciona elemento
   slice[inicio:fim]          // corta slice
   len(slice)                 // tamanho
   slice[1:]                  // do índice 1 até o fim
   slice[:3]                  // do início até índice 3 (exclusivo)

3. PERCORRER ARRAYS/SLICES:
   for i := 0; i < len(arr); i++ { arr[i] }
   for i, valor := range arr { }

💡 EXEMPLO DE SAÍDA ESPERADA:
=== ATIVIDADE 4: ARRAYS E SLICES ===
Temperatura média da semana: 23.89°C
Maior temperatura: 26.1°C, Menor: 21.9°C
Nomes (5 total): 1.Ana 2.Bruno 3.Carlos 4.Diana 5.Eduardo
Primeiros 3 números: [10 20 30]
Últimos 3 números: [70 80 90]
Total de vendas: R$ 9651.30
Vendas acima da média: 3
Números pares: [10 20 30 40 50 60 70 80 90]
Fila atual: [Maria Pedro Ana Carlos]
Matriz 3x3:
1 2 3
4 5 6
7 8 9
Soma da matriz: 45
*/

// Para testar: descomente a linha abaixo
// func main() { exercicioArrays() }
