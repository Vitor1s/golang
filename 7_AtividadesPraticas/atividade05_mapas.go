package main

import "fmt"

/*
🎯 ATIVIDADE 5: MAPAS (MAPS)

📝 OBJETIVOS:
- Criar e manipular mapas (chave-valor)
- Adicionar, remover e verificar elementos
- Iterar sobre mapas
- Resolver problemas práticos com maps

✅ TAREFAS:
*/

func exercicioMapas() {
	fmt.Println("=== ATIVIDADE 5: MAPAS ===")

	// TODO 1: Dicionário de tradução
	// - crie um map[string]string para traduzir português->inglês
	// - adicione: "casa"->"house", "carro"->"car", "livro"->"book", "água"->"water"
	// - teste traduzindo algumas palavras
	// - verifique se uma palavra existe no dicionário antes de traduzir

	// TODO 2: Contador de palavras
	// - frase := "go é uma linguagem go é simples go é rápida"
	// - crie um map[string]int para contar quantas vezes cada palavra aparece
	// - percorra a frase e conte as palavras
	// - imprima o resultado: "go: 3, é: 3, uma: 1, ..."
	// Dica: use strings.Fields() para separar palavras ou faça manualmente

	// TODO 3: Cadastro de idades
	// - crie um map[string]int com nomes e idades
	// - adicione: "João"->25, "Maria"->30, "Pedro"->22, "Ana"->28
	// - encontre a pessoa mais velha e mais nova
	// - calcule a idade média
	// - liste pessoas maiores de 25 anos

	// TODO 4: Estoque de produtos
	// - crie um map[string]int para controlar estoque
	// - produtos: "notebook"->10, "mouse"->25, "teclado"->15, "monitor"->8
	// - simule vendas: venda 3 notebooks, 5 mouses, 2 teclados
	// - simule reposição: adicione 5 monitores, 10 notebooks
	// - imprima estoque final
	// - liste produtos com estoque baixo (menos de 10 unidades)

	// TODO 5: Notas de alunos
	// - crie um map[string][]float64 (cada aluno tem várias notas)
	// - adicione alunos com suas notas:
	//   "João": [7.5, 8.0, 6.5, 9.0]
	//   "Maria": [8.5, 9.0, 8.0, 8.5]
	//   "Pedro": [6.0, 7.0, 6.5, 7.5]
	// - calcule a média de cada aluno
	// - encontre o aluno com maior média

	// TODO 6: Agrupamento por categoria
	// - produtos := []string{"maçã", "banana", "cenoura", "brócolis", "laranja", "alface"}
	// - categorias := []string{"fruta", "fruta", "legume", "legume", "fruta", "legume"}
	// - crie um map[string][]string para agrupar por categoria
	// - resultado esperado: {"fruta": ["maçã", "banana", "laranja"], "legume": ["cenoura", "brócolis", "alface"]}

	// TODO 7: Sistema de login
	// - crie um map[string]string para usuário->senha
	// - adicione alguns usuários: "admin"->"123", "user"->"456", "guest"->"789"
	// - simule tentativas de login:
	//   * tente login("admin", "123") -> sucesso
	//   * tente login("user", "999") -> falha
	//   * tente login("inexistente", "123") -> usuário não existe

	// TODO 8: Análise de frequência de caracteres
	// - texto := "programming"
	// - crie um map[string]int para contar frequência de cada letra
	// - imprima quais letras aparecem mais de uma vez
	// - encontre a letra mais frequente

	// TODO 9: Cache/histórico de buscas
	// - crie um map[string]int para simular cache de buscas
	// - buscas := []string{"golang", "python", "golang", "java", "golang", "javascript", "python"}
	// - para cada busca, incremente o contador
	// - imprima as top 3 buscas mais frequentes

	// ========================================
	// ESCREVA SEU CÓDIGO AQUI EMBAIXO:
	// ========================================

}

/*
🎓 DICAS IMPORTANTES:

1. CRIANDO MAPAS:
   map[tipoChave]tipoValor{}
   make(map[tipoChave]tipoValor)

2. OPERAÇÕES BÁSICAS:
   mapa[chave] = valor                    // adicionar/atualizar
   valor := mapa[chave]                   // acessar
   valor, existe := mapa[chave]           // verificar se existe
   delete(mapa, chave)                    // remover

3. ITERANDO MAPAS:
   for chave, valor := range mapa { }
   for chave := range mapa { }            // só as chaves

4. VERIFICAR EXISTÊNCIA:
   if valor, ok := mapa[chave]; ok {
       // chave existe
   }

💡 EXEMPLO DE SAÍDA ESPERADA:
=== ATIVIDADE 5: MAPAS ===
Tradução de 'casa': house
Contador de palavras: go:3, é:3, uma:1, linguagem:1, simples:1, rápida:1
Pessoa mais velha: Maria (30 anos)
Idade média: 26.25 anos
Estoque de notebook: 17 unidades
Produtos com estoque baixo: monitor (8)
Média do João: 7.75
Aluno com maior média: Maria (8.50)
Login admin: sucesso
Letra mais frequente: r (2 vezes)
Top busca: golang (3 vezes)
*/

// Para testar: descomente a linha abaixo
// func main() { exercicioMapas() }
