package main

import "fmt"

func main() {

	std := make(map[string]string)
	std["name"] = "nahid"
	std["age"] = "20"
	std["tech"] = "CST"

	user := map[string]string{
		"name": "nahid",
		"age":  "20",
	}

	user["city"] = "Tangail"
	fmt.Println(user["name"], user["age"], user["city"])
	delete(user, "city")
	fmt.Println(user)

	for key, value := range std {
		fmt.Println(key, value)
	}

}
