package generics

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type number interface {
	int | float64 | int64
}

type numberAproximacao interface {
	~int | ~float64 | ~int64
}

type myInt int

// Exemplo sem o uso de Generics incluidos na versão 1.18 do GO
func somaInteiro(a, b int64) int64 {
	return a + b
}

func somaFloat(a, b float64) float64 {
	return a + b
}

func Soma[T comparable](a, b T) T {
	if a == b {
		return a
	}

	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func ExemploSemGenerics() {
	fmt.Println(somaInteiro(1, 2))
	fmt.Println(somaFloat(1.2, 2.3))
}

// Em uso de Generics incluidos na versão 1.18 do GO
func soma[T int64 | float64](a, b T) T {
	return a + b
}

func somaComInterface[T numberAproximacao](a, b T) T {
	return a + b
}

func ExemploComGenerics() {

	var x, y myInt = 1, 2

	fmt.Println(somaComInterface(x, y))

	fmt.Println(somaComInterface(1, 2))
	fmt.Println(soma(1.2, 2.3))
}
