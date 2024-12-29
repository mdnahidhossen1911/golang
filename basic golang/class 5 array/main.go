package main

import "fmt"

func main() {
	var name [5]string // type 1
	name[0] = "nahid"
	name[1] = "hossen"
	fmt.Println(name[1])

	stud := [4]string{"Nahid", "nafiz"} //type 2
	fmt.Println(stud)
	fmt.Println(len(stud))
}
