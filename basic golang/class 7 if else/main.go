package main

import "fmt"

func main() {

	i := 1

	/*
		and &&
		or ||
		not  !(a && b)
	*/

	if i > 1 && i == 0 {
		fmt.Println("up")
	} else if i == 1 {
		fmt.Println("Equal")
	} else {
		fmt.Println("down")
	}
}
