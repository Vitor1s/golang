package main

import (
    "encoding/json" // Para converter structs em JSON
    "fmt"           // Para formatação de strings
    "io"            // Para operações de entrada/saída
    "log"           // Para logs do sistema
    "net/http"      // Para criar servidor HTTP
    "sort"          // PACOTE SORT: Para ordenar dados
    "strconv"       // Para converter strings em números
    "strings"       // Para manipulação de strings
    "time"          // Para timestamps
)

// PACOTE SORT + STRUCTS:
// Estrutura de produto para nossa API
type Produto struct {
    ID          int     `json:"id"`
    Nome        string  `json:"nome"`
    Preco       float64 `json:"preco"`
    Categoria   string  `json:"categoria"`
    Estoque     int     `json:"estoque"`
    DataCriacao string  `json:"data_criacao"`
}

// PACOTE SORT + INTERFACE:
// Implementando interface sort.Interface para ordenar por nome
type PorNome []Produto

func (p PorNome) Len() int           { return len(p) }
func (p PorNome) Less(i, j int) bool { return p[i].Nome < p[j].Nome }
func (p PorNome) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// PACOTE SORT + INTERFACE:
// Implementando interface sort.Interface para ordenar por preço
type PorPreco []Produto

func (p PorPreco) Len() int           { return len(p) }
func (p PorPreco) Less(i, j int) bool { return p[i].Preco < p[j].Preco }
func (p PorPreco) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// PACOTE SORT + INTERFACE:
// Implementando interface sort.Interface para ordenar por estoque (decrescente)
type PorEstoque []Produto

func (p PorEstoque) Len() int           { return len(p) }
func (p PorEstoque) Less(i, j int) bool { return p[i].Estoque > p[j].Estoque }
func (p PorEstoque) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// PACOTE SORT + INTERFACE:
// Implementando interface sort.Interface para ordenar por categoria
type PorCategoria []Produto

func (p PorCategoria) Len() int           { return len(p) }
func (p PorCategoria) Less(i, j int) bool { return p[i].Categoria < p[j].Categoria }
func (p PorCategoria) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// PACOTE SORT + STRUCT:
// Banco de dados simulado com produtos
type BancoDados struct {
    produtos []Produto
}

// PACOTE SORT + INICIALIZAÇÃO:
// Inicializa banco com dados de exemplo
func NovoBancoDados() *BancoDados {
    return &BancoDados{
        produtos: []Produto{
            {ID: 1, Nome: "Notebook Dell", Preco: 2500.00, Categoria: "Eletrônicos", Estoque: 15, DataCriacao: "2024-01-15"},
            {ID: 2, Nome: "Mouse Logitech", Preco: 89.90, Categoria: "Acessórios", Estoque: 50, DataCriacao: "2024-01-10"},
            {ID: 3, Nome: "Teclado Mecânico", Preco: 299.90, Categoria: "Acessórios", Estoque: 25, DataCriacao: "2024-01-12"},
            {ID: 4, Nome: "Monitor Samsung", Preco: 1200.00, Categoria: "Eletrônicos", Estoque: 8, DataCriacao: "2024-01-08"},
            {ID: 5, Nome: "Cadeira Gamer", Preco: 899.90, Categoria: "Móveis", Estoque: 12, DataCriacao: "2024-01-05"},
        },
    }
}

// PACOTE SORT + NET/HTTP + ENCODING/JSON:
// Handler para listar produtos com ordenação
func (db *BancoDados) handleListarProdutos(w http.ResponseWriter, r *http.Request) {
    // PACOTE NET/HTTP:
    // Verifica se é método GET
    if r.Method != http.MethodGet {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    // PACOTE STRINGS:
    // Pega parâmetro de ordenação da URL (?sort=nome, ?sort=preco, etc.)
    sortBy := r.URL.Query().Get("sort")
    
    // PACOTE SORT:
    // Cria cópia dos produtos para não modificar o original
    produtos := make([]Produto, len(db.produtos))
    copy(produtos, db.produtos)

    // PACOTE SORT:
    // Aplica ordenação baseada no parâmetro
    switch strings.ToLower(sortBy) {
    case "nome":
        sort.Sort(PorNome(produtos))
    case "preco":
        sort.Sort(PorPreco(produtos))
    case "estoque":
        sort.Sort(PorEstoque(produtos))
    case "categoria":
        sort.Sort(PorCategoria(produtos))
    default:
        // Se não especificar, mantém ordem original
    }

    // PACOTE ENCODING/JSON:
    // Converte produtos ordenados para JSON
    resposta, _ := json.MarshalIndent(produtos, "", "  ")

    // PACOTE NET/HTTP:
    // Define headers da resposta
    w.Header().Set("Content-Type", "application/json")

    // PACOTE IO:
    // Escreve resposta JSON
    io.WriteString(w, string(resposta))
}

// PACOTE SORT + NET/HTTP + ENCODING/JSON:
// Handler para buscar produtos por categoria com ordenação
func (db *BancoDados) handleBuscarPorCategoria(w http.ResponseWriter, r *http.Request) {
    // PACOTE NET/HTTP:
    if r.Method != http.MethodGet {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    // PACOTE STRINGS:
    // Extrai categoria da URL (/categoria/eletronicos)
    partes := strings.Split(r.URL.Path, "/")
    if len(partes) != 3 {
        http.Error(w, "URL inválida", http.StatusBadRequest)
        return
    }

    categoria := partes[2]
    sortBy := r.URL.Query().Get("sort")

    // PACOTE SORT:
    // Filtra produtos por categoria
    var produtosFiltrados []Produto
    for _, produto := range db.produtos {
        if strings.ToLower(produto.Categoria) == strings.ToLower(categoria) {
            produtosFiltrados = append(produtosFiltrados, produto)
        }
    }

    // PACOTE SORT:
    // Aplica ordenação nos produtos filtrados
    switch strings.ToLower(sortBy) {
    case "nome":
        sort.Sort(PorNome(produtosFiltrados))
    case "preco":
        sort.Sort(PorPreco(produtosFiltrados))
    case "estoque":
        sort.Sort(PorEstoque(produtosFiltrados))
    }

    // PACOTE ENCODING/JSON:
    resposta, _ := json.MarshalIndent(produtosFiltrados, "", "  ")

    // PACOTE NET/HTTP:
    w.Header().Set("Content-Type", "application/json")

    // PACOTE IO:
    io.WriteString(w, string(resposta))
}

// PACOTE SORT + NET/HTTP + ENCODING/JSON:
// Handler para produtos com estoque baixo (ordenados por estoque)
func (db *BancoDados) handleEstoqueBaixo(w http.ResponseWriter, r *http.Request) {
    // PACOTE NET/HTTP:
    if r.Method != http.MethodGet {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    // PACOTE STRINGS + STRCONV:
    // Pega limite de estoque da URL (?limite=10)
    limiteStr := r.URL.Query().Get("limite")
    limite := 10 // valor padrão
    if limiteStr != "" {
        if l, err := strconv.Atoi(limiteStr); err == nil {
            limite = l
        }
    }

    // PACOTE SORT:
    // Filtra produtos com estoque baixo
    var produtosEstoqueBaixo []Produto
    for _, produto := range db.produtos {
        if produto.Estoque <= limite {
            produtosEstoqueBaixo = append(produtosEstoqueBaixo, produto)
        }
    }

    // PACOTE SORT:
    // Ordena por estoque (menor primeiro)
    sort.Sort(PorEstoque(produtosEstoqueBaixo))

    // PACOTE ENCODING/JSON:
    resposta, _ := json.MarshalIndent(produtosEstoqueBaixo, "", "  ")

    // PACOTE NET/HTTP:
    w.Header().Set("Content-Type", "application/json")

    // PACOTE IO:
    io.WriteString(w, string(resposta))
}

// PACOTE SORT + NET/HTTP + ENCODING/JSON:
// Handler para criar novo produto
func (db *BancoDados) handleCriarProduto(w http.ResponseWriter, r *http.Request) {
    // PACOTE NET/HTTP:
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
    // Converte JSON para struct
    var novoProduto Produto
    if err := json.Unmarshal(body, &novoProduto); err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }

    // PACOTE SORT:
    // Gera novo ID (maior ID + 1)
    maxID := 0
    for _, produto := range db.produtos {
        if produto.ID > maxID {
            maxID = produto.ID
        }
    }
    novoProduto.ID = maxID + 1
    novoProduto.DataCriacao = time.Now().Format("2006-01-02")

    // PACOTE SORT:
    // Adiciona produto ao banco
    db.produtos = append(db.produtos, novoProduto)

    // PACOTE LOG:
    log.Printf("Novo produto criado: %s (ID: %d)", novoProduto.Nome, novoProduto.ID)

    // PACOTE ENCODING/JSON:
    resposta, _ := json.Marshal(novoProduto)

    // PACOTE NET/HTTP:
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)

    // PACOTE IO:
    io.WriteString(w, string(resposta))
}

func main() {
    // PACOTE SORT:
    // Inicializa banco de dados
    db := NovoBancoDados()

    // PACOTE NET/HTTP:
    // Define rotas da API
    http.HandleFunc("/produtos", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            db.handleListarProdutos(w, r)
        case http.MethodPost:
            db.handleCriarProduto(w, r)
        default:
            http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        }
    })

    http.HandleFunc("/categoria/", db.handleBuscarPorCategoria)
    http.HandleFunc("/estoque-baixo", db.handleEstoqueBaixo)

    // PACOTE LOG:
    // Inicia servidor
    log.Println("API de Produtos iniciada na porta 8080")
    log.Println("Endpoints disponíveis:")
    log.Println("  GET /produtos?sort=nome|preco|estoque|categoria")
    log.Println("  GET /categoria/{categoria}?sort=nome|preco|estoque")
    log.Println("  GET /estoque-baixo?limite=10")
    log.Println("  POST /produtos")
    log.Fatal(http.ListenAndServe(":8080", nil))
}