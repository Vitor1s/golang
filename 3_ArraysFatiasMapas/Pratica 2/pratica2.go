package main 

import "fmt"

func main(){
	praticaslice()
}

func praticaslice(){
	slice := []int{10,15,20,25,30,35,40,45}
	fmt.Println(slice[0:5])
}