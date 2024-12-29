package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//infinity loop
	count := 0
	for {
		fmt.Println(count)
		count++
		if count == 5 {
			break
		} else if count == 3 {
			continue
		}
	}

	// get data
	slice := []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		fmt.Println("index:", index, "\nvalue", value)
	}

	name := "nahid"
	for index, char := range name {
		fmt.Printf("index %d char %c\n", index, char)
	}

}
