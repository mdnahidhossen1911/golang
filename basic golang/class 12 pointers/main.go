package main

/// pointers stors memory address

import "fmt"

func main() {

	num := 23
	var pointers *int = &num
	fmt.Println(pointers)

	name := "nahid"
	pit := &name
	fmt.Println(pit)

	var pnt *int // deafult value of pointers nil
	fmt.Println(pnt)

	///modifyValueByRefarence
	value := 10
	modifyValueByRefarence(&value) //func
	fmt.Println(value)
}

func modifyValueByRefarence(value *int) {
	*value = *value + 20
}
