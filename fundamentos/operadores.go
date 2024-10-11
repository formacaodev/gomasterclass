package fundamentos

import "fmt"

func aritmeticos() {
	fmt.Println("Soma:", 10+2)
	fmt.Println("Subtração:", 10-2)
	fmt.Println("Multiplicação:", 10*2)
	fmt.Println("Divisão:", 10/2)
	fmt.Println("Módulo:", 10%2)
}

func relacionais() {
	fmt.Println("Igual:", 10 == 2)
	fmt.Println("Diferente:", 10 != 2)
	fmt.Println("Maior:", 10 > 2)
	fmt.Println("Maior ou igual:", 10 >= 2)
	fmt.Println("Menor:", 10 < 2)
	fmt.Println("Menor ou igual:", 10 <= 2)
}

func logicos() {
	fmt.Println("E:", true && false)
	fmt.Println("OU:", true || false)
	fmt.Println("NÃO:", !true)
}
