package mian

import "fmt"


func main() {
	pratica4()
}

// //TODO.  Na pratica4 vamos explorar os tipos os tipos quando definiidos nao podem ser mudados, exemplo e quando definimos var x int  x = 10, dentro do bloco nao podemos mudar esse tipo chamando ele novamente mas mudadando x = "vitor" ou 10.5  pois um e string e o outro é float. 
func pratica4(){
// Mostrando na pratica como isso ocorre 
	var x int = 10
	fmt.Println(x)
	x = "vitor"
	fmt.Println(x)	

	// TODO. Como pode ver eu defini a variavel *x* coimo int, mas depois tentei atribuir uma string a ela, o que causou um erro de compilação. Isso demonstra que uma vez que um tipo é definido, ele não pode ser alterado para outro tipo dentro do mesmo escopo. E pra isso dar certo voce tera que criar outra variavel com outro nome, ou seja, se eu quisesse atribuir uma string a uma variavel, eu teria que criar outra variavel do tipo string, como por exemplo:

	var y string = "vitor"
	fmt.Println(y)

// 	//TODO. Uma obs pra voce descobri o tipo da variavel é só uasa o %T depois do %g que usamoos no printf para cahmar a variavel, por exemplo:
	fmt.Printf("O tipo de x é %T e o tipo de y é %T.\n", x, y)
 }