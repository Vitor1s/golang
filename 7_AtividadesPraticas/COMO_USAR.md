# 🎯 INSTRUÇÕES PARA AS ATIVIDADES

## 📋 Como Usar Este Material

### 1. **Ordem Recomendada**
Execute as atividades na sequência para melhor aprendizado:

1. `atividade01_tipos.go` - Conceitos básicos
2. `atividade02_variaveis_correto.go` - Variáveis e escopo  
3. `atividade03_controle.go` - If/else e loops
4. `atividade04_arrays.go` - Arrays e slices
5. `atividade05_mapas.go` - Estruturas chave-valor
6. `atividade06_structs.go` - Estruturas de dados
7. `atividade07_funcoes.go` - Funções avançadas
8. `atividade08_arquivos.go` - Manipulação de arquivos
9. `atividade09_projeto_completo.go` - Projeto integrado
10. `atividade10_sistema_completo.go` - Desafio final

### 2. **Como Testar Cada Atividade**

Para testar uma atividade específica:

```bash
# Navegue até a pasta
cd /Users/vitorandrade/Projetos/meusProjetos/golang/7_AtividadesPraticas

# Execute a atividade desejada
go run atividade01_tipos.go
```

**IMPORTANTE:** Para as atividades funcionarem, você precisa descomentar a linha `func main()` no final de cada arquivo e implementar as soluções.

### 3. **Estrutura dos Arquivos**

Cada arquivo tem:
- **Objetivos claros** do que praticar
- **TODOs numerados** com tarefas específicas
- **Dicas importantes** sobre conceitos
- **Exemplo de saída esperada**
- **Espaço para seu código**

### 4. **Dicas Para Resolver**

#### ✅ **Se estiver com dificuldade:**
1. Releia suas anotações no arquivo `Contexto da linguagem.txt`
2. Consulte seus códigos anteriores nas pastas `1_PrimeirosPassos`, `2_EstruturasEmGo`, etc.
3. Use o arquivo `exemplos_solucoes.go` como referência (só se necessário!)

#### ✅ **Abordagem recomendada:**
1. Leia todos os TODOs antes de começar
2. Implemente um TODO por vez
3. Teste frequentemente com `go run`
4. Não passe para o próximo até ter certeza

#### ✅ **Para debugar problemas:**
```bash
# Se der erro de compilação
go run arquivo.go

# Para ver detalhes do erro
go build arquivo.go
```

### 5. **Exemplos de Como Começar**

#### **Atividade 1 (Tipos):**
```go
func main() {
    fmt.Println("=== ATIVIDADE 1: TIPOS E OPERADORES ===")
    
    // TODO 1: Crie variáveis usando inferência
    nome := "Seu Nome Aqui"
    idade := 25
    // ... continue implementando
}
```

#### **Atividade 4 (Arrays):**
```go
func exercicioArrays() {
    fmt.Println("=== ATIVIDADE 4: ARRAYS E SLICES ===")
    
    // TODO 1: Array de temperaturas
    var temperaturas [7]float64
    temperaturas[0] = 23.5
    // ... continue preenchendo
}
```

### 6. **Verificação de Progresso**

Marque conforme for completando:

- [ ] Atividade 1: Tipos básicos ✅
- [ ] Atividade 2: Variáveis e escopo ✅  
- [ ] Atividade 3: Estruturas de controle ✅
- [ ] Atividade 4: Arrays e slices ✅
- [ ] Atividade 5: Mapas ✅
- [ ] Atividade 6: Structs ✅
- [ ] Atividade 7: Funções avançadas ✅
- [ ] Atividade 8: Manipulação de arquivos ✅
- [ ] Atividade 9: Projeto integrado ✅
- [ ] Atividade 10: Sistema completo ✅

### 7. **Recursos Adicionais**

#### **Para consulta rápida:**
- `Contexto da linguagem.txt` - Suas anotações
- `1_PrimeirosPassos/` - Fundamentos
- `2_EstruturasEmGo/` - Controle de fluxo
- `3_ArraysFatiasMapas/` - Estruturas de dados
- `4_EstruturasInterfaces/` - Structs
- `5_FunçõesemGo/` - Funções
- `6_PacotesEmGo/` - Pacotes e arquivos

#### **Se precisar de ajuda:**
1. Consulte a documentação oficial: `go doc fmt.Println`
2. Use `exemplos_solucoes.go` para ver exemplos
3. Revise seus códigos anteriores

### 8. **Desafios Extras**

Após completar todas as atividades:

1. **Combine conceitos:** Misture funcionalidades de diferentes atividades
2. **Otimize código:** Refatore para tornar mais eficiente
3. **Adicione features:** Implemente funcionalidades extras
4. **Crie testes:** Teste diferentes cenários e casos extremos

### 9. **Comandos Úteis**

```bash
# Formatar código
go fmt arquivo.go

# Verificar sintaxe sem executar
go build arquivo.go

# Executar com debugging
go run -x arquivo.go

# Ver documentação de uma função
go doc fmt.Printf
```

---

## 🎉 **Parabéns!**

Você criou um conjunto completo de atividades que cobrem **TUDO** que aprendeu em Go:

✅ **Conceitos básicos** → Tipos, variáveis, operadores  
✅ **Estruturas de controle** → If/else, loops  
✅ **Estruturas de dados** → Arrays, slices, maps, structs  
✅ **Funções** → Parâmetros, retornos, recursão  
✅ **Pacotes** → Manipulação de arquivos, I/O  
✅ **Projetos completos** → Sistemas integrados  

Agora é hora de **praticar, praticar e praticar!** 🚀

**Lembre-se:** A programação se aprende programando. Não tenha medo de errar, é assim que se aprende!

---

**Boa sorte com suas atividades! 💪**
