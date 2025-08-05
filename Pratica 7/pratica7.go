package main

import( "fmt"
 		"os"

)

func main() {
	pratica7()

}

//TODO. Nessa pratica vamos entender a diferenças do pacote "fmt" 1 = func Print /Printf/ Println  2 = func Sprint / Sprintf / Sprintln 3 = func Fprint / Fprintf / Fprintln

func pratica7(){ 

    // 1. Print, Printf, Println
    fmt.Println("1. Print, Printf, Println:")

    // Print: escreve tudo junto, não pula linha no final
    fmt.Print("Oi, eu sou o Print. ")
    fmt.Print("Eu não pulo linha!\n")

    // Println: escreve e pula linha no final
    fmt.Println("Oi, eu sou o Println. Eu pulo linha no final!")

    // Printf: escreve formatando, como um molde
    nome := "Maria"
    idade := 7
    fmt.Printf("Oi, meu nome é %s e eu tenho %d anos.\n", nome, idade)

    fmt.Println("\n2. Sprint, Sprintf, Sprintln:")

    // 2. Sprint, Sprintf, Sprintln (eles criam strings ao invés de mostrar na tela)
    // Sprint: junta tudo numa string
    texto := fmt.Sprint("Eu sou o Sprint. ", "Eu junto tudo numa frase.")
    fmt.Println(texto)

    // Sprintln: junta tudo numa string e coloca um \n no final
    texto2 := fmt.Sprintln("Eu sou o Sprintln.", "Eu também pulo linha no final!")
    fmt.Print(texto2) // Usando Print porque texto2 já tem \n

    // Sprintf: faz igual ao Printf, mas devolve uma string
    texto3 := fmt.Sprintf("Meu nome é %s e tenho %d anos.", nome, idade)
    fmt.Println(texto3)

    fmt.Println("\n3. Fprint, Fprintf, Fprintln:")

    // 3. Fprint, Fprintf, Fprintln (escrevem em outro lugar, como um arquivo ou tela especial)
    // Aqui vamos escrever na tela do computador (os.Stdout)
    fmt.Fprint(os.Stdout, "Eu sou o Fprint. Escrevo onde você mandar!\n")
    fmt.Fprintln(os.Stdout, "Eu sou o Fprintln. Também pulo linha!")
    fmt.Fprintf(os.Stdout, "Eu sou o Fprintf. Meu nome é %s!\n", nome)

    // Você também pode escrever em arquivos, mas aqui só mostramos na tela.


}


// Print, Printf, Println: mostram na tela.
// Sprint, Sprintf, Sprintln: criam textos (strings) para você usar depois.
// Fprint, Fprintf, Fprintln: escrevem onde você quiser (tela, arquivo, etc).
