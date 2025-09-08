package main

import "fmt"

/*
🎯 ATIVIDADE 1: TIPOS E OPERADORES BÁSICOS

📝 OBJETIVOS:
- Trabalhar com todos os tipos básicos (string, int, float, bool)
- Usar operadores lógicos (&&, ||, !)
- Praticar inferência de tipos
- Trabalhar com strings (len, concatenação, indexação)

✅ TAREFAS:
*/

func main() {
	fmt.Println("=== ATIVIDADE 1: TIPOS E OPERADORES ===")

	// TODO 1: Crie variáveis usando inferência de tipos:
	// - nome (seu nome)
	// - idade (sua idade)
	// - altura (sua altura)
	// - estudante (true se estuda)

	// TODO 2: Imprima uma apresentação usando concatenação de strings
	// Exemplo: "Olá, eu sou João, tenho 25 anos..."

	// TODO 3: Crie duas variáveis booleanas:
	// - temComputador = true
	// - temInternet = true
	// Use operadores lógicos para verificar:
	// a) Se pode programar (precisa dos dois)
	// b) Se tem pelo menos um recurso (|| - ou)
	// c) Negue o resultado de "pode programar" (!)

	// TODO 4: Trabalhe com strings:
	// - conte quantos caracteres tem seu nome (len)
	// - mostre a primeira letra do seu nome (indexação [0])
	// - mostre a última letra (dica: use len() - 1)
	// - crie uma frase juntando nome + sobrenome

	// TODO 5: Teste números:
	// - crie duas notas (float64): nota1 = 8.5, nota2 = 9.2
	// - calcule a média
	// - verifique se a média é >= 7.0 (aprovado)

	// TODO 6: Teste com ASCII:
	// - imprima o código ASCII da letra 'A' (use fmt.Println("A"[0]))
	// - compare com o que você aprendeu sobre Unicode/ASCII

	// ========================================
	// ESCREVA SEU CÓDIGO AQUI EMBAIXO:
	// ========================================

}

/*
🎓 DICAS:
- Lembre-se: var nome = "valor" (inferência automática)
- Operadores: && (e), || (ou), ! (não)
- String[index] retorna código ASCII, não a letra
- Para converter ASCII para letra use: string(codigo)

💡 EXEMPLO DE SAÍDA ESPERADA:
=== ATIVIDADE 1: TIPOS E OPERADORES ===
Olá, eu sou Vitor, tenho 25 anos, meço 1.80m e sou estudante: true
Posso programar: true
Tenho pelo menos um recurso: true
NÃO posso programar: false
Meu nome tem 5 caracteres
Primeira letra: V (código 86)
Última letra: r (código 114)
Nome completo: Vitor Andrade
Média das notas: 8.85
Aprovado: true
*/
