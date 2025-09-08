package main

import "fmt"

/*
🎯 ATIVIDADE 3: ESTRUTURAS DE CONTROLE

📝 OBJETIVOS:
- Praticar if/else em diferentes situações
- Dominar loops for em suas variações
- Combinar condicionais com loops
- Trabalhar com operadores de comparação

✅ TAREFAS:
*/

func controle() {
	fmt.Println("=== ATIVIDADE 3: ESTRUTURAS DE CONTROLE ===")

	// TODO 1: Classificador de idade
	// - crie uma variável idade := 17
	// - use if/else para classificar:
	//   * menor que 13: "Criança"
	//   * 13 a 17: "Adolescente"
	//   * 18 a 59: "Adulto"
	//   * 60 ou mais: "Idoso"

	// TODO 2: Verificador de nota
	// - crie nota := 8.5
	// - use if/else if para mostrar conceito:
	//   * 9.0-10.0: "Excelente"
	//   * 7.0-8.9: "Bom"
	//   * 5.0-6.9: "Regular"
	//   * 0.0-4.9: "Insuficiente"

	// TODO 3: Loop básico - conte de 1 a 10
	// Use for com: for i := 1; i <= 10; i++

	// TODO 4: Loop while - conte de 50 até 45 (regressivo)
	// Use: contador := 50; for contador >= 45 { ... }

	// TODO 5: Tabuada personalizada
	// - peça um número (ex: numero := 7)
	// - imprima a tabuada desse número (1x7=7, 2x7=14, etc.)

	// TODO 6: Números pares e ímpares
	// - use for de 1 a 20
	// - dentro do loop, use if para verificar se é par ou ímpar
	// - imprima "X é par" ou "X é ímpar"
	// Dica: use i%2 == 0 para verificar se é par

	// TODO 7: Contador de positivos, negativos e zeros
	// - crie um slice: numeros := []int{-3, 0, 5, -1, 8, 0, -7, 2}
	// - use for range para percorrer
	// - conte quantos são positivos, negativos e zeros
	// - imprima o resultado

	// TODO 8: Encontrar o maior número
	// - use o mesmo slice do TODO 7
	// - encontre o maior número usando loop
	// - dica: comece com maior := numeros[0]

	// TODO 9: Simulador de caixa eletrônico
	// - saldo := 1000.0
	// - operacoes := []string{"saque", "deposito", "saque", "deposito"}
	// - valores := []float64{200.0, 150.0, 50.0, 300.0}
	// - use for range para processar cada operação
	// - se saque: diminui do saldo (mas só se tiver saldo suficiente)
	// - se depósito: aumenta o saldo
	// - imprima o saldo após cada operação

	// TODO 10: Detector de sequência
	// - numeros := []int{1, 2, 3, 4, 5, 10, 11, 12}
	// - encontre sequências de números consecutivos
	// - imprima algo como: "Sequência encontrada: 1,2,3,4,5" e "Sequência encontrada: 10,11,12"

	// ========================================
	// ESCREVA SEU CÓDIGO AQUI EMBAIXO:
	// ========================================

}

/*
🎓 DICAS IMPORTANTES:

1. OPERADORES DE COMPARAÇÃO:
   == (igual), != (diferente), < (menor), > (maior), <= (menor igual), >= (maior igual)

2. TIPOS DE FOR:
   for i := 0; i < 10; i++ { }     // C-style loop
   for condicao { }                 // while loop
   for index, valor := range slice { }  // range loop

3. IF MÚLTIPLOS:
   if condicao1 {
   } else if condicao2 {
   } else {
   }

4. RANGE:
   - for i, v := range slice: i=índice, v=valor
   - for _, v := range slice: só o valor (ignora índice)
   - for i := range slice: só o índice

💡 EXEMPLO DE SAÍDA ESPERADA:
=== ATIVIDADE 3: ESTRUTURAS DE CONTROLE ===
Idade 17: Adolescente
Nota 8.5: Bom
Contagem: 1 2 3 4 5 6 7 8 9 10
Regressivo: 50 49 48 47 46 45
Tabuada do 7: 1x7=7, 2x7=14, 3x7=21...
15 é ímpar, 16 é par, 17 é ímpar...
Positivos: 3, Negativos: 3, Zeros: 2
Maior número: 8
Saldo após saque: 800.00
Sequência encontrada: 1,2,3,4,5
*/
