// package main

// import( 
// 	"fmt"
// 	"sort"
// )

// type Dados struct{
// 	Nome string
// 	  int

// }

// type ParaNome[]Dados // Aqui criamos a estrutura ParaNime e abrimos a interface []Dados para que possamos usar o sort.Slice
// func (ps ParaNome) Len()int{
// 	return len(ps)
// }
// func (ps ParaNome) Less(i, j int)bool{ // Aqui criamos a função Less que compara os elementos da lista vai ver se i é menor que j

// 	return ps[i].Nome < ps[j].Nome

// }
// func (ps ParaNome) Swap(i, j int){ // O que Swap faz é trocar os elementos da lista

// 	ps[i], ps[j] = ps[j], ps[i]

// }

// func main(){
// 	 criancas := []Dados{
// 		{"A", 10},
// 		{"B", 8},
// 		{"C", 12},
// 	 }
// 	 sort.Sort(ParaNome(criancas))
// 	 fmt.Println(criancas)
// 	 sort.Sort(ParaIdade(criancas))
// 	 fmt.Println(ParaIdade(criancas))
// }

// type ParaIdade[]Dados
// func (ps ParaIdade) Len()int{
// 	return len(ps)
// }
// func (ps ParaIdade) Less(i, j int)bool{
// 	return ps[i].Idade > ps[j].Idade
// }
// func (ps ParaIdade) Swap(i, j int){
// 	ps[i], ps[j] = ps[j], ps[i]
// }