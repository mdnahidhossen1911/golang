package main

import (
	"fmt"
	"strconv"
)

func main() {

	//int
	number := 123
	fmt.Printf("%T\n", number)

	//int -> float
	f_num := float64(number)
	fmt.Printf("%T\n", f_num)

	//int -> string
	s_num := string(number)
	fmt.Printf("%T\n", s_num)

	//string ->int
	i_num, _ := strconv.Atoi(s_num)
	fmt.Printf("%T\n", i_num)

	//str -> float 64
	sf_num, _ := strconv.ParseFloat(s_num, 64)
	fmt.Printf("%T\n", sf_num)

}
