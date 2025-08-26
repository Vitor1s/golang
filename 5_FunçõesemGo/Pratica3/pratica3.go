package main

import "fmt"

// defer agenda funções para rodarem no fim da execução da função atual, em ordem inversa.
// panic interrompe o programa de forma abrupta.
// recover captura o panic e permite que o programa continue rodando.


func dia1(){
	fmt.Println("Domingo")
}

func dia2(){
	fmt.Println("Segunda")
}

func dia3(){
	fmt.Println("Terça") 
}

func main(){
	defer dia1()
	defer dia2()
	defer dia3()
}

/// Aqui vamos ver o panic que pode vir por erro de tempo de execução ou por erro do desenvolvedor e etc. e tambem vamos ver o recover que pode ser usado para recuperar o programa de um panic. podendo escalonar o programa de forma mais organizada.


// E qual a utilidade de usar o panic? Ela é usada para parar o programa e mostrar o erro, e o recover é usado para recuperar o programa de um panic.

func main(){
	defer func(){
		x := recover()
		fmt.Println(x)
	}()
	panic("Erro")
}
