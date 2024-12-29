package main

import "fmt"

func main() {

	user := []int{1, 2, 3, 4} // array update version no need array size
	fmt.Println(user)
	fmt.Println(len(user))
	user = append(user, 5) //slice add item

	// fmt.Println("//////////")
	// number := make([]int, 3, 4) // lenth 3  , capacity 4 // item 0 0 0
	// fmt.Println(number)
	// fmt.Println(cap(number)) // capacity
	// fmt.Println(len(number)) //lenth

}
