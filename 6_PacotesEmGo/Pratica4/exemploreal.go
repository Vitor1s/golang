package main

import (
    "fmt"
    "os"
    "time"
)

// // PACOTE TIME + ERRO PERSONALIZADO:
// // Erro personalizado para quando arquivo não é encontrado
// type ErroArquivoNaoEncontrado struct {
//     NomeArquivo string    // Nome do arquivo que não foi encontrado
//     HoraErro    time.Time // PACOTE TIME: Quando o erro aconteceu
// }

// // PACOTE FMT + INTERFACE ERROR:
// // Implementando o método Error() da interface 'error'
// func (e ErroArquivoNaoEncontrado) Error() string {
//     return fmt.Sprintf("ERRO REAL: Arquivo '%s' não encontrado às %s", 
//         e.NomeArquivo, e.HoraErro.Format("15:04:05"))
// }

// // PACOTE OS + ERRO REAL:
// // Função que tenta abrir um arquivo (pode dar erro REAL)
// func tentarAbrirArquivo(nomeArquivo string) error {
//     // PACOTE OS:
//     // Tentando abrir arquivo - se não existir, dá erro REAL
//     _, err := os.Open(nomeArquivo)
//     if err != nil {
//         // PACOTE TIME:
//         // Criando erro personalizado com informações úteis
//         return ErroArquivoNaoEncontrado{
//             NomeArquivo: nomeArquivo,
//             HoraErro:    time.Now(),
//         }
//     }
//     return nil // Sem erro
// }

// func main() {
//     // Testando com arquivo que NÃO EXISTE (erro REAL)
//     err := tentarAbrirArquivo("arquivo_que_nao_existe.txt")
//     if err != nil {
//         fmt.Println(err) // Vai mostrar: "ERRO REAL: Arquivo 'arquivo_que_nao_existe.txt' não encontrado às 14:30:25"
//     }
    
//     // Testando com arquivo que EXISTE (sem erro)
//     err = tentarAbrirArquivo("pratica4.go") // Este arquivo existe
//     if err != nil {
//         fmt.Println(err)
//     } else {
//         fmt.Println("Arquivo encontrado com sucesso!")
//     }
// }