package main

import "fmt"

func main() {
	i := 5
	switch i {
	case 1:
		fmt.Print("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("no")
		fmt.Println(i)
	}
}
