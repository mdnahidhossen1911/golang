package main

import "fmt"

func main() {
	name("Nahid")
	user("Nahid", 21)
	number(23, 33)
	fmt.Println(sum(23, 54))
	fmt.Println(math(23, 54, 76))
}

func name(name string) { //type 1
	fmt.Println(name)
}

func user(name string, age int) { //type 2
	fmt.Println(name)
	fmt.Println(age)
}

func number(user1, user2 int) { // type 3
	fmt.Println(user1)
	fmt.Println(user2)
}

func sum(a, b int) int { //int return type 1
	return a + b
}

func math(a, b, c int) (result int) { // return type 2
	result = a + b + c
	return
}
