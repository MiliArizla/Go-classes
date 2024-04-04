package main

import "fmt"

func main(){

	hobbies := make(map[string]string)
	hobbies["Julia"] = "to code"
	hobbies["Ana"] = "to play with dogs"
	hobbies["Gabi"]= "to read"

	for name, hobby := range hobbies {
		fmt.Println( name,"'s hobby is:", hobby)
	}
}