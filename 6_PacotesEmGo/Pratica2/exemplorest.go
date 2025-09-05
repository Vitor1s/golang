package main

import (
    "encoding/json" // Para converter structs em JSON
    "fmt"           // Para formatação de strings
    "io"            // Para operações de entrada/saída
    "log"           // Para logs do sistema
    "net/http"      // Para criar servidor HTTP
    "os"            // Para manipulação de arquivos
    "strconv"       // Para converter strings em números
    "strings"       // Para manipulação de strings
    "time"          // Para timestamps
	"bytes"         // Para manipulação de bytes
) 

// PACOTE OS + STRUCTS:
// Definindo estrutura de dados para usuário
type Usuario struct {
    ID       int    `json:"id"`       // Tag para JSON
    Nome     string `json:"nome"`     // Campo será "nome" no JSON
    Email    string `json:"email"`    // Campo será "email" no JSON
    CriadoEm string `json:"criado_em"` // Campo será "criado_em" no JSON
}

// PACOTE BYTES + IO:
// Simulando banco de dados em arquivo JSON
type BancoDados struct {
    arquivo string
    usuarios []Usuario
}

// PACOTE OS + IO:
// Carrega dados do arquivo JSON
func (db *BancoDados) carregarDados() error {
    // os.Open() - Abre arquivo para leitura
    arquivo, err := os.Open(db.arquivo)
    if err != nil {
        // Se arquivo não existe, cria um novo
        if os.IsNotExist(err) {
            db.usuarios = []Usuario{}
            return db.salvarDados()
        }
        return err
    }
    defer arquivo.Close() // Garante que arquivo será fechado

    // PACOTE BYTES + IO:
    // Lê todo conteúdo do arquivo
    buffer := new(bytes.Buffer)
    buffer.ReadFrom(arquivo)
    
    // PACOTE ENCODING/JSON:
    // Converte JSON do buffer para slice de usuários
    return json.Unmarshal(buffer.Bytes(), &db.usuarios)
}

// PACOTE OS + IO:
// Salva dados no arquivo JSON
func (db *BancoDados) salvarDados() error {
    // PACOTE ENCODING/JSON:
    // Converte slice de usuários para JSON
    dados, err := json.MarshalIndent(db.usuarios, "", "  ")
    if err != nil {
        return err
    }

    // os.Create() - Cria/sobrescreve arquivo
    arquivo, err := os.Create(db.arquivo)
    if err != nil {
        return err
    }
    defer arquivo.Close()

    // PACOTE IO:
    // Escreve JSON no arquivo
    _, err = io.WriteString(arquivo, string(dados))
    return err
}

// PACOTE STRINGS + LOG:
// Função para criar novo usuário
func (db *BancoDados) criarUsuario(nome, email string) (*Usuario, error) {
    // PACOTE STRINGS:
    // Validação básica - verifica se email contém "@"
    if !strings.Contains(email, "@") {
        return nil, fmt.Errorf("email inválido")
    }

    // Cria novo usuário
    novoUsuario := Usuario{
        ID:       len(db.usuarios) + 1,
        Nome:     nome,
        Email:    email,
        CriadoEm: time.Now().Format("2006-01-02 15:04:05"),
    }

    db.usuarios = append(db.usuarios, novoUsuario)

    // PACOTE LOG:
    // Registra criação do usuário
    log.Printf("Novo usuário criado: %s (%s)", nome, email)

    return &novoUsuario, db.salvarDados()
}

// PACOTE NET/HTTP + ENCODING/JSON:
// Handler para criar usuário via POST
func (db *BancoDados) handleCriarUsuario(w http.ResponseWriter, r *http.Request) {
    // PACOTE NET/HTTP:
    // Verifica se é método POST
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    // PACOTE IO:
    // Lê corpo da requisição
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Erro ao ler requisição", http.StatusBadRequest)
        return
    }

    // PACOTE ENCODING/JSON:
    // Converte JSON da requisição para struct
    var dados struct {
        Nome  string `json:"nome"`
        Email string `json:"email"`
    }
    
    if err := json.Unmarshal(body, &dados); err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }

    // Cria usuário
    usuario, err := db.criarUsuario(dados.Nome, dados.Email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // PACOTE ENCODING/JSON:
    // Converte usuário criado para JSON
    resposta, _ := json.Marshal(usuario)
    
    // PACOTE NET/HTTP:
    // Define headers da resposta
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    
    // PACOTE IO:
    // Escreve resposta JSON
    io.WriteString(w, string(resposta))
}

// PACOTE NET/HTTP + ENCODING/JSON:
// Handler para listar usuários via GET
func (db *BancoDados) handleListarUsuarios(w http.ResponseWriter, r *http.Request) {
    // PACOTE NET/HTTP:
    // Verifica se é método GET
    if r.Method != http.MethodGet {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    // PACOTE ENCODING/JSON:
    // Converte lista de usuários para JSON
    resposta, _ := json.MarshalIndent(db.usuarios, "", "  ")
    
    // PACOTE NET/HTTP:
    // Define headers da resposta
    w.Header().Set("Content-Type", "application/json")
    
    // PACOTE IO:
    // Escreve resposta JSON
    io.WriteString(w, string(resposta))
}

// PACOTE NET/HTTP + STRINGS + STRCONV:
// Handler para buscar usuário por ID via GET
func (db *BancoDados) handleBuscarUsuario(w http.ResponseWriter, r *http.Request) {
    // PACOTE NET/HTTP:
    // Verifica se é método GET
    if r.Method != http.MethodGet {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    // PACOTE STRINGS:
    // Extrai ID da URL (ex: /usuarios/123)
    partes := strings.Split(r.URL.Path, "/")
    if len(partes) != 3 {
        http.Error(w, "URL inválida", http.StatusBadRequest)
        return
    }

    // PACOTE STRCONV:
    // Converte string do ID para número
    id, err := strconv.Atoi(partes[2])
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    // Busca usuário
    for _, usuario := range db.usuarios {
        if usuario.ID == id {
            // PACOTE ENCODING/JSON:
            // Converte usuário para JSON
            resposta, _ := json.Marshal(usuario)
            
            // PACOTE NET/HTTP:
            w.Header().Set("Content-Type", "application/json")
            
            // PACOTE IO:
            io.WriteString(w, string(resposta))
            return
        }
    }

    // PACOTE NET/HTTP:
    // Usuário não encontrado
    http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

func main() {
    // PACOTE OS:
    // Inicializa banco de dados com arquivo
    db := &BancoDados{arquivo: "usuarios.json"}
    
    // PACOTE IO + LOG:
    // Carrega dados existentes
    if err := db.carregarDados(); err != nil {
        log.Printf("Erro ao carregar dados: %v", err)
    }

    // PACOTE NET/HTTP:
    // Define rotas da API
    http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            db.handleListarUsuarios(w, r)
        case http.MethodPost:
            db.handleCriarUsuario(w, r)
        default:
            http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        }
    })

    http.HandleFunc("/usuarios/", db.handleBuscarUsuario)

    // PACOTE LOG:
    // Inicia servidor
    log.Println("API iniciada na porta 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}