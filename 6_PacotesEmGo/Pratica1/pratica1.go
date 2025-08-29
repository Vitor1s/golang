package main

import "fmt"
import "strings" //. esse pacote traz funções para manipulação de strings podendo procurar strings dentro de outras strings, converter para maiúsculas ou minúsculas, dividir strings em partes, entre outras funcionalidades.

// // Definindo uma struct chamada retangulo
// // Structs agrupam dados relacionados em um único tipo
// // Imagine que você tem um retângulo, como uma folha de papel.

// aqui vamos usar a funçao contains que é uma funçao que verifica se uma string contém outra string.

func main() {
	texto := "Hello, World! Welcome to Go programming."
	contem := strings.Contains(texto, "Vitor")
	fmt.Println(contem)
}