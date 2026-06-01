package functions

import "fmt"

func Incrementar(p *int) {
	fmt.Println(p)
	*p++
}

func Decrementar(p *int) {
	fmt.Println(p)
	*p--
}

type Person struct {
	Name string
	Age  int
}

// Se eu nao passar o *Person o p.Age++ so vai funcionar dentro do escopo da funcao, se eu colocar o *Person vai alterar o valor da variavel fora da funcao diretamente no endereco da memoria.
func (p *Person) Aniversario() {
	p.Age++
}
