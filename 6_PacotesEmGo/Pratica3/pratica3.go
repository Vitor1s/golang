package main

import (
	"fmt"
	"os"
	"path/filepath" // Para manipulação de caminhos de arquivos
)

// PACOTE PATH/FILEPATH + OS:
// filepath.Walk() percorre todos os arquivos e pastas de um diretório
// É como um "explorador de arquivos" que visita cada item
func main() {
	// filepath.Walk(".", callback) - Percorre o diretório atual (".")
	// "." significa "diretório onde estou agora" (Pratica3/)
	//Detalhe importante é que ele procura a pasta usando o mesmo padrao de terminal, e como se ele saisse da pasta e visse a pasta proncipal 6_PacotesEmGo e dps achasse a pasta Pratica1
	filepath.Walk("../Pratica1", func(path string, info os.FileInfo, err error) error {
		// PACOTE OS:
		// Se houve erro ao acessar arquivo/pasta
		if err != nil {
			return err
		}

		// PACOTE OS:
		// info.IsDir() retorna true se for pasta, false se for arquivo
		// path = caminho completo (ex: "./Pratica1/pratica1.go")
		// info.IsDir() = é pasta? (true/false)
		fmt.Println(path, info.IsDir())
		return nil
	})
}
