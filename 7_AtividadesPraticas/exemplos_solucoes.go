package main

import "fmt"

/*
🎯 EXEMPLOS DE SOLUÇÕES

Este arquivo contém exemplos de como resolver algumas das atividades.
Use como referência se ficar travado em algum exercício!

⚠️ IMPORTANTE: Tente resolver sozinho primeiro! Só consulte se realmente precisar.
*/

// ==========================================
// EXEMPLO: SOLUÇÃO DA ATIVIDADE 1 (TIPOS)
// ==========================================

func exemploAtividade1() {
	fmt.Println("=== EXEMPLO ATIVIDADE 1: TIPOS ===")

	// TODO 1: Variáveis com inferência
	nome := "Vitor"
	idade := 25
	altura := 1.75
	estudante := true

	// TODO 2: Apresentação com concatenação
	apresentacao := fmt.Sprintf("Olá, eu sou %s, tenho %d anos, meço %.2fm e sou estudante: %t",
		nome, idade, altura, estudante)
	fmt.Println(apresentacao)

	// TODO 3: Operadores lógicos
	temComputador := true
	temInternet := true

	podeProgramar := temComputador && temInternet
	temAlgumRecurso := temComputador || temInternet
	naoPodeProgramar := !podeProgramar

	fmt.Println("Pode programar:", podeProgramar)
	fmt.Println("Tem algum recurso:", temAlgumRecurso)
	fmt.Println("NÃO pode programar:", naoPodeProgramar)

	// TODO 4: Manipulação de strings
	qtdCaracteres := len(nome)
	primeiraLetra := nome[0]
	ultimaLetra := nome[len(nome)-1]

	fmt.Println("Nome tem", qtdCaracteres, "caracteres")
	fmt.Printf("Primeira letra: %c (código %d)\n", primeiraLetra, primeiraLetra)
	fmt.Printf("Última letra: %c (código %d)\n", ultimaLetra, ultimaLetra)

	// TODO 5: Cálculos com notas
	nota1 := 8.5
	nota2 := 9.2
	media := (nota1 + nota2) / 2
	aprovado := media >= 7.0

	fmt.Printf("Média: %.2f\n", media)
	fmt.Println("Aprovado:", aprovado)
}

// ==========================================
// EXEMPLO: SOLUÇÃO PARCIAL ATIVIDADE 4 (ARRAYS)
// ==========================================

func exemploAtividade4() {
	fmt.Println("=== EXEMPLO ATIVIDADE 4: ARRAYS ===")

	// TODO 1: Array de temperaturas
	temperaturas := [7]float64{23.5, 25.0, 22.8, 24.2, 26.1, 21.9, 23.7}

	// Calcular média
	var soma float64
	for _, temp := range temperaturas {
		soma += temp
	}
	media := soma / float64(len(temperaturas))

	// Encontrar maior e menor
	maior := temperaturas[0]
	menor := temperaturas[0]
	for _, temp := range temperaturas {
		if temp > maior {
			maior = temp
		}
		if temp < menor {
			menor = temp
		}
	}

	fmt.Printf("Média da semana: %.2f°C\n", media)
	fmt.Printf("Maior: %.1f°C, Menor: %.1f°C\n", maior, menor)

	// TODO 2: Slice de nomes
	nomes := []string{"Ana", "Bruno", "Carlos"}
	nomes = append(nomes, "Diana", "Eduardo")

	fmt.Printf("Total de nomes: %d\n", len(nomes))
	for i, nome := range nomes {
		fmt.Printf("%d. %s\n", i+1, nome)
	}

	// TODO 3: Manipulação de slice
	numeros := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	primeiros3 := numeros[0:3]
	ultimos3 := numeros[len(numeros)-3:]
	meio := numeros[3:6]

	fmt.Println("Primeiros 3:", primeiros3)
	fmt.Println("Últimos 3:", ultimos3)
	fmt.Println("Meio:", meio)
}

// ==========================================
// EXEMPLO: FUNÇÃO COM MÚLTIPLOS RETORNOS
// ==========================================

func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}
	return a / b, nil
}

func exemploFuncaoComErro() {
	fmt.Println("=== EXEMPLO: FUNÇÃO COM ERRO ===")

	// Teste com valores válidos
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", resultado)
	}

	// Teste com erro
	resultado, err = dividir(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Resultado: %.2f\n", resultado)
	}
}

// ==========================================
// EXEMPLO: STRUCT SIMPLES
// ==========================================

type Pessoa struct {
	nome  string
	idade int
	email string
}

func exemploStruct() {
	fmt.Println("=== EXEMPLO: STRUCT ===")

	// Criando pessoa
	p1 := Pessoa{
		nome:  "João",
		idade: 30,
		email: "joao@email.com",
	}

	// Forma alternativa
	p2 := Pessoa{"Maria", 25, "maria@email.com"}

	fmt.Printf("Pessoa 1: %s, %d anos, %s\n", p1.nome, p1.idade, p1.email)
	fmt.Printf("Pessoa 2: %s, %d anos, %s\n", p2.nome, p2.idade, p2.email)

	// Comparar idades
	if p1.idade > p2.idade {
		fmt.Printf("%s é mais velha\n", p1.nome)
	} else {
		fmt.Printf("%s é mais velha\n", p2.nome)
	}
}

// ==========================================
// EXEMPLO: MAP BÁSICO
// ==========================================

func exemploMap() {
	fmt.Println("=== EXEMPLO: MAP ===")

	// Criar map
	idades := make(map[string]int)

	// Adicionar dados
	idades["João"] = 30
	idades["Maria"] = 25
	idades["Pedro"] = 35

	// Acessar dados
	if idade, existe := idades["João"]; existe {
		fmt.Printf("João tem %d anos\n", idade)
	}

	// Iterar sobre map
	fmt.Println("Todas as idades:")
	for nome, idade := range idades {
		fmt.Printf("- %s: %d anos\n", nome, idade)
	}

	// Encontrar pessoa mais velha
	var maisVelha string
	var maiorIdade int
	for nome, idade := range idades {
		if idade > maiorIdade {
			maiorIdade = idade
			maisVelha = nome
		}
	}
	fmt.Printf("Pessoa mais velha: %s (%d anos)\n", maisVelha, maiorIdade)
}

// ==========================================
// FUNÇÃO PRINCIPAL PARA TESTAR EXEMPLOS
// ==========================================

func exemplosSolucoes() {
	fmt.Println("🎓 EXEMPLOS DE SOLUÇÕES")
	fmt.Println("==================================================")

	exemploAtividade1()
	fmt.Println()

	exemploAtividade4()
	fmt.Println()

	exemploFuncaoComErro()
	fmt.Println()

	exemploStruct()
	fmt.Println()

	exemploMap()
}

/*
🎓 DICAS PARA USAR ESTES EXEMPLOS:

1. ANALISE O CÓDIGO:
   - Entenda cada linha antes de copiar
   - Veja como os conceitos se conectam
   - Note os padrões de nomenclatura

2. ADAPTE PARA SUAS ATIVIDADES:
   - Use como base, não como cópia exata
   - Modifique para atender os requisitos específicos
   - Adicione suas próprias melhorias

3. PRATIQUE VARIAÇÕES:
   - Mude os valores de teste
   - Adicione novas funcionalidades
   - Combine diferentes conceitos

4. APRENDA COM ERROS:
   - Execute o código e veja o que acontece
   - Faça modificações intencionais para ver erros
   - Entenda as mensagens de erro

💡 LEMBRE-SE:
O objetivo não é apenas fazer funcionar, mas ENTENDER como funciona!
*/

// Para testar os exemplos: descomente a linha abaixo
// func main() { exemplosSolucoes() }
