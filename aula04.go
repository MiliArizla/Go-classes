package main

// Array

// type Employee struct {
// 	name   string
// 	age    int
// 	role   string
// 	salary float64
// }

// // Colocar uma structure dentro de um array

// func main() {

// 	employee1 := Employee{name: "Julia", age: 22, role: "developer", salary: 8500.50}
// 	employee2 := Employee{name: "Gabriela", age: 25, role: "support", salary: 8525.40}
// 	employee3 := Employee{name: "Alexia", age: 26, role: "scrum master", salary: 8250.98}

// 	var employees [3]Employee

// 	employees[0] = employee1
// 	employees[1] = employee2
// 	employees[2] = employee3

// 	fmt.Println(employees[1])
// }

// Slice

// func main() {
// 	hello := "Hello"

// 	slice := []string{hello}
// 	slice = append(slice, "Julia", "Gabriela", "Alexia")

// 	secondName := slice[1]

// 	fmt.Print(slice)
// 	fmt.Printf("\n%s, %s", slice[0], secondName)

// }

// Map
func main(){

	// votes:= map[string]int{"Yan": 4, "Jones" : 2, "White": 2}
	// votos := make(map[string]int) // Create an empty map
	// votos = {"Maria":2, "Ana": 4}

	contacts := make(map[string]string)
	contacts["Julia"] = "54 9999-9999"
	contacts["Ana"] = "54 9999-9998"
	contacts["Gabi"]= "54 9999-9997"

	// for name, phone := range contacts {
	// 	fmt.Println("O número da ", name, "é ", phone)
	// }
}