package main

import "fmt"

//  ===== Bem Vindo =====
// func main() {
//   var nome string

//   fmt.Print("Qual seu nome? ")
//   fmt.Scan(&nome)
//   fmt.Println("Bem Vindo(a) ", nome,"!")

//  ===== Soma =====
// func main() {
// 	var number int
// 	five := 5

// 	fmt.Println("Insira um número:")
// 	fmt.Scan(&number)
// 	fmt.Println("A soma de 5 +",number , "é",number + five)
// }

// ===== Condicionais If Else =====
// func main() {
// 	var number int

// 	fmt.Println("Insira um número:")
// 	fmt.Scan(&number)
// 	if (number > 0){
// 		fmt.Println("Positivo!")
// 	} else if (number < 0) {
// 		fmt.Println("Negativo!")
// 	} else if (number == 0){
// 		fmt.Println("Zero!")
// 	}
// }

// ===== Condicionais Switch Case =====
func main() {
	var number int
	number = 3

	fmt.Println("Insira um número:")
	num, err := fmt.Scan(&number)
	fmt.Println(num, err)
	fmt.Println("your number is", number)
	switch{
	case number > 0:
		fmt.Println("Positivo!")
	case number < 0:
		fmt.Println("Negativo!")
	case number == 0:
		fmt.Println("Zero!")
		}
	}

// ===== Loops FOR =====
// func main() {
// 	for n:= 10; n>0; n-- {
// 		fmt.Println("Contagem Regressiva", n)
// 	}
// 	fmt.Println("Arrasou!")
// }
