package main

import (
	"fmt"
	// "time" // PACOTE TIME: Armazena quando o erro aconteceu
	"container/list" // PACOTE LIST: Armazena uma lista de elementos
)

// // PACOTE TIME + ERROS PERSONALIZADOS:
// // Aqui vamos ter uma ideia de como funcionar o pacote de error do go

// // PACOTE TIME + STRUCT:
// // Definindo um tipo de erro personalizado
// // Para criar erros customizados, você precisa implementar a interface 'error'
// type MeuErro struct {
//     hora   time.Time // PACOTE TIME: Armazena quando o erro aconteceu
//     minuto string    // Mensagem personalizada do erro
// }

// // PACOTE FMT + INTERFACE ERROR:
// // Implementando o método Error() da interface 'error'
// // Go automaticamente chama este método quando você imprime o erro
// func (e *MeuErro) Error() string {
//     // PACOTE FMT:
//     // Sprintf formata uma string sem imprimir (retorna string)
//     // %s = placeholder para string (e.minuto)
//     return fmt.Sprintf("Erro: %s", e.minuto)
// }

// // PACOTE TIME + FMT:
// func main() {
//     // PACOTE TIME:
//     // time.Now() retorna o momento atual
//     // &MeuErro{} cria um ponteiro para a struct
//     meuErro := &MeuErro{hora: time.Now(), minuto: "00"}

//     // PACOTE FMT:
//     // fmt.Println() automaticamente chama o método Error() do erro
//     // Por isso aparece "Erro: 00" em vez de mostrar a struct completa
//     fmt.Println(meuErro)
// }

// PACOTE TIME + ERROS PERSONALIZADOS:
// Abaixo mais um exemplo de como funciona o pacote de error do go

// PACOTE TIME + STRUCT:
// Outro exemplo de erro personalizado com mais informações
// type MyError struct {
// 	When time.Time // PACOTE TIME: Quando o erro aconteceu
// 	What string    // O que causou o erro
// }

// // PACOTE FMT + INTERFACE ERROR:
// // Implementando o método Error() da interface 'error'
// func (e MyError) Error() string {
// 	// PACOTE FMT:
// 	// %v = placeholder genérico (funciona com qualquer tipo)
// 	// %v para time.Time mostra data formatada
// 	// %v para string mostra o texto
// 	return fmt.Sprintf("%v: %v", e.When, e.What)
// }

// // PACOTE TIME + ERRO:
// // Função que retorna um erro personalizado
// func oops() error {
// 	// PACOTE TIME:
// 	// time.Date() cria uma data específica
// 	// time.Date(ano, mês, dia, hora, minuto, segundo, nanosegundo, timezone)
// 	return MyError{
// 		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC), // 15 de março de 1989, 22:30
// 		"O Arquivo desapareceu",                        // Mensagem do erro
// 	}
// }

// // PACOTE FMT + TRATAMENTO DE ERRO:
// func main() {
// 	// Chamando função que pode retornar erro
// 	// if err := oops(); err != nil { ... } é o padrão Go para tratar erros
// 	if err := oops(); err != nil {
// 		// PACOTE FMT:
// 		// fmt.Println(err) automaticamente chama o método Error()
// 		// Por isso aparece: "1989-03-15 22:30:00 +0000 UTC: the file system has gone away"
// 		fmt.Println(err)
// 	}
// }

func main() {
	// Create a new list and put some numbers in it.

    // Vou explicar o que entendi desse codigo criamos a lista e atribuimos ao nome lista o list.New() que vem do pacote container/list  e depois criamos dois elementos atribuidos a func lista e4 e e1 que vem do mesmo pacote. e usamos eles para definir a ordem da lista  onde pushBack é o final e PushFront o iicio  depois dentro da lista criamos mais dois lementos mas esses de inserçao IsertBefore que coloca o tres antes do 4 e InsertAfter que coloca o 2 depois do 1. Depois imprimmimos essa lista com for que olha a lista  no for criamos uma variavel olhar onde ele olha o front que seria o 1 da lista e Next que é o proximo elemento da lista/ A unica coisa que nao entendi foi esse Value?? Pq ele ali??
	lista := list.New()
	e4 := lista.PushBack(4)
	e1 := lista.PushFront(1)
	lista.InsertBefore(3, e4)
	lista.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for olhar := lista.Front(); olhar != nil; olhar = olhar.Next() {
		fmt.Println(olhar.Value)
	}

}
