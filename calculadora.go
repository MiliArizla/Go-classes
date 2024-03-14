package main

import "fmt"

func main() {

	var number1 int
	var number2 int
	var operacao string

	fmt.Println("Insira o primeiro número:")
	fmt.Scan(&number1)

	fmt.Println("Insira o segundo número:")
	fmt.Scan(&number2)

	fmt.Println("Insira o tipo de  operação: (+, -, /, *)")
	fmt.Scan(&operacao)

	switch {
	case operacao == "+":
		fmt.Println("O resultado da soma de", number1, operacao, number2,"é",number1 + number2)
	case operacao == "-":
		fmt.Println("O resultado da diferença de", number1, operacao, number2,"é",number1 - number2)
	case operacao == "/":
		fmt.Println("O resultado da divisão de", number1, operacao, number2,"é",number1 / number2)
	case operacao == "*":
		fmt.Println("O resultado da multiplicação de", number1, operacao, number2, "é",number1 * number2)
	default:
		fmt.Println("Operação inválida")
	}
}