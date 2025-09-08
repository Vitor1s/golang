package main

import "fmt"

/*
🎯 ATIVIDADE 2: VARIÁVEIS E ESCOPO

📝 OBJETIVOS:
- Praticar diferentes formas de declarar variáveis
- Entender blocos e escopo
- Trabalhar com reatribuição de valores
- Comparar var, := e declaração separada

✅ TAREFAS:
*/

// TODO 1: Declare uma variável global chamada "linguagem" com valor "Go"
// (fora da função principal, descomente a linha abaixo)
// var linguagem = "Go"

func principal1() {
	fmt.Println("=== ATIVIDADE 2: VARIÁVEIS E ESCOPO ===")

	// TODO 2: Declare variáveis de 3 formas diferentes:
	// Forma 1: var explicita com tipo
	// var produto string = "Notebook"

	// Forma 2: var com inferência
	// var preco = 2500.99

	// Forma 3: declaração curta
	// quantidade := 3

	// TODO 3: Teste reatribuição:
	// - mude o valor de produto para "Mouse"
	// - mude o preço para 45.90
	// - aumente a quantidade em 2 unidades

	// TODO 4: Crie um bloco (usando chaves {}) e dentro dele:
	// - declare uma variável local "desconto := 0.1"
	// - calcule precoComDesconto = preco * (1 - desconto)
	// - imprima o resultado
	// - tente imprimir "desconto" fora do bloco (vai dar erro - comente essa linha)

	// TODO 5: Trabalhe com múltiplas declarações:
	// var (
	//     cliente = "João"
	//     cidade  = "São Paulo"
	//     cep     = "01234-567"
	// )

	// TODO 6: Use a variável global "linguagem":
	// - imprima "Programando em: " + linguagem

	// TODO 7: Calcule o total da compra:
	// total := float64(quantidade) * preco
	// (note a conversão de int para float64)

	// ========================================
	// ESCREVA SEU CÓDIGO AQUI EMBAIXO:
	// ========================================

}

/*
🎓 EXPLICAÇÕES:

1. FORMAS DE DECLARAR VARIÁVEIS:
   var nome tipo = valor    // mais explícita
   var nome = valor         // com inferência
   nome := valor           // forma curta (só dentro de funções)

2. ESCOPO:
   - Variáveis dentro de {} só existem nesse bloco
   - Variáveis globais (fora de funções) são acessíveis em qualquer lugar
   - Variáveis da função só existem dentro da função

3. REATRIBUIÇÃO:
   - Depois de declarada, use apenas: nome = novoValor
   - Não use := novamente (isso criaria nova variável)

💡 EXEMPLO DE SAÍDA ESPERADA:
=== ATIVIDADE 2: VARIÁVEIS E ESCOPO ===
Produto: Mouse, Preço: R$ 45.90, Quantidade: 5
Preço com desconto: R$ 41.31
Cliente: João de São Paulo, CEP: 01234-567
Programando em: Go
Total da compra: R$ 229.50
*/

// Para testar, descomente a linha abaixo:
// func main() { principal() }
