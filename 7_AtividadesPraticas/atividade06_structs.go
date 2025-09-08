package main

import "fmt"

/*
🎯 ATIVIDADE 6: STRUCTS (ESTRUTURAS)

📝 OBJETIVOS:
- Criar e usar structs para organizar dados
- Trabalhar com structs aninhadas
- Implementar métodos (functions em structs)
- Resolver problemas do mundo real com structs

✅ TAREFAS:
*/

func exercicioStructs() {
	fmt.Println("=== ATIVIDADE 6: STRUCTS ===")

	// TODO 1: Struct básica - Pessoa
	// - crie uma struct Pessoa com: nome string, idade int, email string
	// - crie 3 pessoas diferentes
	// - imprima os dados de cada pessoa formatado
	// - encontre a pessoa mais velha

	// TODO 2: Struct de Produto para loja
	// - crie struct Produto com: nome string, preco float64, categoria string, estoque int
	// - crie alguns produtos: notebook, mouse, teclado
	// - calcule o valor total do estoque (preco * quantidade para todos produtos)
	// - liste produtos com estoque baixo (menos de 5 unidades)

	// TODO 3: Struct para Endereço e Pessoa completa
	// - crie struct Endereco: rua string, numero int, cidade string, cep string
	// - modifique struct Pessoa para incluir um campo endereco Endereco
	// - crie uma pessoa com endereço completo
	// - imprima os dados formatados: "João, 25 anos, mora na Rua A, 123, São Paulo"

	// TODO 4: Sistema de Conta Bancária
	// - crie struct ContaBancaria: titular string, numero int, saldo float64
	// - crie funções para:
	//   * depositar(conta *ContaBancaria, valor float64)
	//   * sacar(conta *ContaBancaria, valor float64) bool // retorna se conseguiu sacar
	//   * consultarSaldo(conta ContaBancaria)
	// - teste todas as operações

	// TODO 5: Escola com Alunos e Notas
	// - crie struct Aluno: nome string, idade int, notas []float64
	// - crie struct Turma: professor string, materia string, alunos []Aluno
	// - crie uma turma com 3 alunos
	// - calcule a média da turma
	// - encontre o aluno com maior média

	// TODO 6: Sistema de Biblioteca
	// - crie struct Livro: titulo string, autor string, ano int, disponivel bool
	// - crie struct Biblioteca: nome string, livros []Livro
	// - implemente funções:
	//   * adicionarLivro(bib *Biblioteca, livro Livro)
	//   * emprestarLivro(bib *Biblioteca, titulo string) bool
	//   * devolverLivro(bib *Biblioteca, titulo string)
	//   * listarDisponiveis(bib Biblioteca)

	// TODO 7: E-commerce simples
	// - crie struct ItemCarrinho: produto Produto, quantidade int
	// - crie struct Carrinho: itens []ItemCarrinho, cliente string
	// - implemente funções:
	//   * adicionarItem(carrinho *Carrinho, produto Produto, qtd int)
	//   * calcularTotal(carrinho Carrinho) float64
	//   * listarItens(carrinho Carrinho)

	// TODO 8: Sistema de Funcionários
	// - crie struct Funcionario: nome string, cargo string, salario float64, departamento string
	// - crie struct Empresa: nome string, funcionarios []Funcionario
	// - implemente:
	//   * adicionarFuncionario(empresa *Empresa, func Funcionario)
	//   * calcularFolhaPagamento(empresa Empresa) float64
	//   * listarPorDepartamento(empresa Empresa, dept string)
	//   * funcionarioMaiorSalario(empresa Empresa) Funcionario

	// TODO 9: Jogo simples - Personagem RPG
	// - crie struct Personagem: nome string, nivel int, vida int, ataque int, defesa int
	// - crie struct Equipamento: nome string, tipoEquip string, bonusAtaque int, bonusDefesa int
	// - implemente:
	//   * equipar(personagem *Personagem, equip Equipamento)
	//   * atacar(atacante Personagem, defensor *Personagem)
	//   * mostrarStatus(personagem Personagem)

	// ========================================
	// ESCREVA SEU CÓDIGO AQUI EMBAIXO:
	// ========================================

}

/*
🎓 DICAS IMPORTANTES:

1. DEFININDO STRUCTS:
   type NomeStruct struct {
       campo1 tipo1
       campo2 tipo2
   }

2. CRIANDO INSTÂNCIAS:
   var p Pessoa
   p = Pessoa{nome: "João", idade: 25}
   p = Pessoa{"João", 25, "joao@email.com"}  // ordem dos campos

3. ACESSANDO CAMPOS:
   p.nome = "Novo nome"
   idade := p.idade

4. PONTEIROS EM STRUCTS:
   func modificar(p *Pessoa) {
       p.nome = "Modificado"  // Go faz automaticamente (*p).nome
   }

5. STRUCTS ANINHADAS:
   type Pessoa struct {
       nome string
       endereco Endereco
   }

💡 EXEMPLO DE SAÍDA ESPERADA:
=== ATIVIDADE 6: STRUCTS ===
Pessoa: João, 25 anos, joao@email.com
Pessoa mais velha: Maria (30 anos)
Valor total do estoque: R$ 15.750,00
Produtos com estoque baixo: mouse (3 unidades)
João, 25 anos, mora na Rua das Flores, 123, São Paulo
Depósito realizado. Saldo atual: R$ 1.500,00
Saque negado. Saldo insuficiente.
Média da turma: 7.83
Melhor aluno: Ana (8.5)
Livro 'Dom Casmurro' emprestado com sucesso
Livros disponíveis: O Cortiço, Senhora
Total do carrinho: R$ 2.847,50
Folha de pagamento: R$ 25.000,00
Funcionários de TI: João (Programador), Maria (Analista)
Maior salário: Pedro (R$ 12.000,00)
João atacou Maria causando 15 de dano!
Maria: Nível 5, Vida: 85/100, Ataque: 25, Defesa: 20
*/

// Para testar: descomente a linha abaixo
// func main() { exercicioStructs() }
