package main

import (
	"fmt"
	// Descomente os imports conforme precisar:
	// "os"
	// "io"
	// "log"
	// "bytes"
	// "bufio"
	// "strings"
	// "encoding/json"
)

/*
🎯 ATIVIDADE 8: PACOTES E MANIPULAÇÃO DE ARQUIVOS

📝 OBJETIVOS:
- Trabalhar com pacotes io, os, log, bytes
- Criar, ler, escrever e manipular arquivos
- Processar dados em JSON
- Implementar logging de atividades
- Trabalhar com entrada/saída de dados

✅ TAREFAS:
*/

func exercicioPacotes() {
	fmt.Println("=== ATIVIDADE 8: PACOTES E ARQUIVOS ===")

	// TODO 1: Criação e escrita de arquivo básico
	// - use os.Create() para criar "dados.txt"
	// - escreva algumas linhas no arquivo
	// - use io.WriteString() ou file.WriteString()
	// - feche o arquivo com defer

	// TODO 2: Leitura de arquivo com buffer
	// - leia o arquivo "dados.txt" criado no TODO 1
	// - use bytes.Buffer para ler todo conteúdo
	// - imprima o conteúdo lido na tela
	// - trate possíveis erros com log.Fatal()

	// TODO 3: Sistema de log personalizado
	// - crie arquivo "sistema.log"
	// - use log.SetOutput() para direcionar logs para arquivo
	// - registre diferentes tipos de eventos:
	//   * log.Println("Sistema iniciado")
	//   * log.Printf("Usuário %s fez login", "admin")
	//   * log.Println("Operação realizada com sucesso")

	// TODO 4: Processamento de dados CSV simples
	// - crie arquivo "vendas.csv" com dados:
	//   "Produto,Quantidade,Preco\nNotebook,2,2500.00\nMouse,5,45.90\n"
	// - leia linha por linha
	// - calcule total de vendas
	// - conte quantos produtos diferentes

	// TODO 5: Sistema de configuração JSON
	// - crie uma struct Config com: nome, versao, debug bool, portas []int
	// - crie um arquivo "config.json" com configurações
	// - leia o JSON e carregue na struct
	// - modifique algum valor e salve de volta

	// TODO 6: Monitor de arquivos
	// - verifique se arquivo existe com os.Stat()
	// - obtenha informações: tamanho, data modificação
	// - liste arquivos de um diretório com os.ReadDir()
	// - filtre apenas arquivos .txt ou .go

	// TODO 7: Backup e restore simples
	// - crie função que copia um arquivo para outro (backup)
	// - use io.Copy() para copiar conteúdo
	// - crie função restore que restaura do backup
	// - teste fazendo backup de um arquivo e restaurando

	// TODO 8: Sistema de cache em arquivo
	// - implemente cache que salva dados em "cache.json"
	// - funções: salvarCache(chave, valor), lerCache(chave), limparCache()
	// - use map[string]interface{} para armazenar diferentes tipos

	// TODO 9: Gerador de relatório
	// - leia dados de um slice de structs (ex: vendas)
	// - gere relatório em formato texto
	// - salve em "relatorio.txt" com formatação:
	//   ===== RELATÓRIO DE VENDAS =====
	//   Produto: X, Qtd: Y, Total: Z
	//   ...
	//   TOTAL GERAL: R$ XXXX

	// TODO 10: Sistema de import/export
	// - export: salve slice de structs em JSON
	// - import: carregue JSON de volta para slice
	// - valide dados durante import
	// - crie logs de auditoria para operações

	// ========================================
	// ESCREVA SEU CÓDIGO AQUI EMBAIXO:
	// ========================================

}

/*
🎓 DICAS IMPORTANTES:

1. MANIPULAÇÃO DE ARQUIVOS:
   file, err := os.Create("arquivo.txt")
   if err != nil { log.Fatal(err) }
   defer file.Close()

2. LEITURA COM BUFFER:
   arquivo, _ := os.Open("arquivo.txt")
   buffer := new(bytes.Buffer)
   buffer.ReadFrom(arquivo)
   conteudo := buffer.String()

3. JSON EM GO:
   // Marshal: struct -> JSON
   data, _ := json.Marshal(estrutura)

   // Unmarshal: JSON -> struct
   json.Unmarshal(data, &estrutura)

4. VERIFICAR ARQUIVOS:
   if _, err := os.Stat("arquivo.txt"); os.IsNotExist(err) {
       // arquivo não existe
   }

5. LISTAR DIRETÓRIO:
   entries, _ := os.ReadDir(".")
   for _, entry := range entries {
       if !entry.IsDir() { // é arquivo
           fmt.Println(entry.Name())
       }
   }

💡 EXEMPLO DE SAÍDA ESPERADA:
=== ATIVIDADE 8: PACOTES E ARQUIVOS ===
Arquivo 'dados.txt' criado com sucesso
Conteúdo lido: "Linha 1\nLinha 2\nLinha 3"
Logs salvos em 'sistema.log'
CSV processado: 2 produtos, total R$ 5229.50
Config carregada: {nome: "MeuApp", versao: "1.0", debug: true}
Arquivo 'dados.txt': 156 bytes, modificado hoje
Arquivos .go encontrados: main.go, config.go
Backup realizado: dados.txt -> dados.bak
Cache salvo: {"usuario": "admin", "tema": "dark"}
Relatório gerado em 'relatorio.txt'
Export realizado: 15 registros salvos
Import realizado: 15 registros carregados
*/

// Para testar: descomente a linha abaixo
// func main() { exercicioPacotes() }
