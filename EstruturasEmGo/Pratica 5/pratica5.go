package main

import "fmt"

func main(){
	// praticaloopi()
	// praticalooph()
	// praticaloopwhile()
	praticaloopinfinito()
	praticaloopcontinue()
}

// TODO. Aqui vamos ver o loop com Inicialização, Condição e Pós.
// Inicialização: executa 1 vez antes de começar o loop.
// Condição: verificada antes de cada repetição; se for falsa, o loop para.
// Pós: executa no final de cada repetição.
func praticaloopi() {
	for i := 0;       // Inicialização
		i < 10;       // Condição
		i++ {         // Pós
		fmt.Println(i)
	}
}

// TODO. Aqui vamos ver Loop Repetição Hierárquica
func praticalooph() {
	// Loop externo: controla as horas do dia
	// Começamos com a variável `hora` valendo 0
	// Enquanto `hora` for menor que 24, executa o bloco
	// Ao final de cada repetição, incrementa `hora` em +1
	for hora := 0; hora < 24; hora++ {
		// Imprime o valor atual de `hora` na tela
		fmt.Println("Hora:", hora)

		// Loop interno: controla os minutos da hora atual
		// Começa com `min` igual a 0
		// Enquanto `min` for menor que 60, executa o bloco
		// Ao final de cada repetição, incrementa `min` em +1
		for min := 0; min < 60; min++ {
			// Imprime o valor atual de `min` na tela
			fmt.Println("Minuto:", min)
		}
	}
}



//TODO. Aqui vamos ver o loop while em Go while(enquanto)
func praticaloopwhile(){
	 i := 0

	for ; i < 5;  {
		fmt.Println("i é menor que 5")
		i++
	} 


}


//TODO. O ultimo vamos ver sobre Loop infinito Break e Continue

func praticaloopinfinito() {
   i := 0 

   for { if i < 5 {
	i++
	  fmt.Println("i é menor que 5")
   } else {
       break
   }
}
}

func praticaloopcontinue() {

	for x := 0; x < 20; x++ {
		if x == 5 {
			continue
		}
		fmt.Println(x)
	}
}