package structs

import (
	"encoding/json"
	"fmt"
)

// Tag são usados nas structs para definir metadados, exemplo o json quando estiver usando o marshal `json:"---"`
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PersonIntenatial struct {
	Person
	Country string `json:"country"`
}

var Person1 = Person{
	Name: "Maura",
	Age:  27,
}

var PersonIntenatial1 = PersonIntenatial{
	Person:  Person{Name: "Reinaldo", Age: 25},
	Country: "Brazil",
}

func (p Person) Exibe() {
	fmt.Println(p)
}

func (p *Person) AlteraNome(newName string) {
	p.Name = newName
	fmt.Printf("Seu novo nome é: %s\n", p.Name)
}

func ConvertToJSON(p PersonIntenatial) {
	clienteJson, err := json.Marshal(p)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(clienteJson))
}

func ConvertJSONToStruct(jsonString string) {
	var p PersonIntenatial
	json.Unmarshal([]byte(jsonString), &p)
	fmt.Println(p)
}
