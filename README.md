Seu material está bom, mas tem alguns erros de sintaxe do Go e dá para deixá-lo mais organizado e próximo da documentação oficial.

# 📚 Anotações de Go (Golang)

## Variáveis

### Declaração de variáveis

```go
var nome string = "Reinaldo"
var idade int = 25
var altura float64 = 1.73
var estaAcordado bool = true
```

### Inferência de tipo

O Go consegue identificar automaticamente o tipo da variável através do valor atribuído.

```go
nome := "Reinaldo"
idade := 25
altura := 1.73
estaAcordado := true
```

Nesse caso:

```go
count := 0
```

O compilador entende que `count` é do tipo `int`.

### Declaração sem valor inicial (Zero Value)

Quando uma variável é declarada sem receber um valor, o Go atribui automaticamente um valor padrão chamado **Zero Value**.

```go
var nome string
var idade int
var altura float64
var ativo bool
```

Valores atribuídos automaticamente:

| Tipo    | Zero Value |
| ------- | ---------- |
| string  | `""`       |
| int     | `0`        |
| float64 | `0`        |
| bool    | `false`    |

Exemplo:

```go
var nome string

fmt.Println(nome) // ""
```

---

# Funções

## Função com parâmetros e retorno

Recebe dois inteiros e retorna a soma.

```go
func soma(param1 int, param2 int) int {
	return param1 + param2
}
```

Uso:

```go
resultado := soma(10, 20)
```

---

## Função com múltiplos retornos

Recebe duas strings e retorna ambas.

```go
func retornaNomes(nome1 string, nome2 string) (string, string) {
	return nome1, nome2
}
```

Uso:

```go
primeiro, segundo := retornaNomes("João", "Maria")
```

---

## Função sem retorno

Executa uma ação, mas não retorna nenhum valor.

```go
func printaNome(nome string, sobrenome string) {
	fmt.Printf("Seu nome é %s %s\n", nome, sobrenome)
}
```

Uso:

```go
printaNome("Reinaldo", "Silva")
```

---

## Retorno nomeado (Named Return)

O valor retornado recebe um nome na assinatura da função.

```go
func somaRetornoNomeado(param1 int, param2 int) (soma int) {
	soma = param1 + param2
	return
}
```

É equivalente a:

```go
func soma(param1 int, param2 int) int {
	return param1 + param2
}
```

Embora exista, o retorno nomeado não é muito utilizado no dia a dia.

---

# Ponteiros

Ponteiros armazenam o endereço de memória de uma variável.

## Operador `&`

O operador `&` retorna o endereço de memória de uma variável.

```go
count := 0

fmt.Println(&count)
```

Saída:

```text
0xc0000120a0
```

---

## Operador `*`

O operador `*` acessa o valor armazenado no endereço apontado pelo ponteiro.

```go
count := 0

var ponteiro *int = &count

fmt.Println(*ponteiro)
```

Saída:

```text
0
```

---

## Modificando o valor através do ponteiro

```go
count := 0

ponteiro := &count

*ponteiro = 10

fmt.Println(count)
```

Saída:

```text
10
```

O valor original foi alterado porque o ponteiro aponta diretamente para a posição de memória da variável.

---

# Resumo Rápido

### Declarar variável

```go
var nome string = "Reinaldo"
nome := "Reinaldo"
```

### Criar função

```go
func soma(a int, b int) int {
	return a + b
}
```

### Obter endereço de memória

```go
&variavel
```

### Obter valor do ponteiro

```go
*ponteiro
```

### Criar ponteiro

```go
ponteiro := &variavel
```

### Alterar valor usando ponteiro

```go
*ponteiro = novoValor
```

---

💡 **Observação:** Diferente de JavaScript, Go não possui `undefined`. Toda variável sempre possui um valor válido (Zero Value). Já o conceito mais próximo de `null` é `nil`, utilizado para ponteiros, slices, maps, interfaces, channels, funções e alguns outros tipos de referência.

**Nome dado as threads do GO** -> `User Land` ou `Green Threads` e so ocupa 2k de memoria cada thread.