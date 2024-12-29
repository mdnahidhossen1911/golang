package main

import "fmt"

func main() {

	number := 3.4444
	name := "nahid"
	school := "habhit"
	age := 34
	fmt.Println(name)
	fmt.Printf("number: %.2f\n", number)                                               //brack line \n and float 2f
	fmt.Printf("type: %T\n", name)                                                     // name type string
	fmt.Printf("number: %2f name: %s age: %d school: %s\n", number, name, age, school) //number %d // float %2f // string %s
	fmt.Println("naem:", name, " school:", school)

}
