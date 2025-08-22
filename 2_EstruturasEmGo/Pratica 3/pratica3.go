package main

import "fmt"

//TODO. Aqui vamos ver sobre os Switch e como ele pode melhorar a funcionalidade do if

func main() {
	semSwitch()
	comSwitch()
}

// TODO. Na função abaixo mostro como fica sem o switch, ele facilita o processo onde tem varios if
func semSwitch() {
	for i := 0; i <= 5; i++ {
		if i == 0 {
			fmt.Println("ZERO")
		} else if i == 1 {
			fmt.Println("UM")
		} else if i == 2 {
			fmt.Println("DOIS")
		} else if i == 3 {
			fmt.Println("TRES")
		} else if i == 4 {
			fmt.Println("QUATRO")
		} else if i == 5 {
			fmt.Println("CINCO")
		}
	}
}

//TODO. Como pode ser visto quando nao usado o switch, o processo 
func comSwitch() {
	for i := 0; i <= 5; i++ {
		switch i {
		case 0:
			fmt.Println("ZERO")
		case 1:
			fmt.Println("UM")
		case 2:
			fmt.Println("DOIS")
		case 3:
			fmt.Println("TRES")
		case 4:
			fmt.Println("QUATRO")
		case 5:
			fmt.Println("CINCO")
		default:
			fmt.Println("NÃO É UM NÚMERO VÁLIDO")
		}
	}
}
