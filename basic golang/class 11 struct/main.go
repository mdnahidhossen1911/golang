package main

import "fmt"

//create custom datatype
type User struct {
	name   string
	id     int
	city   string
	number int
}

//student
type Info struct {
	name  string
	fName string
	mName string
}

type Contect struct {
	number int
	gmail  string
}

type Address struct {
	country     string
	city        string
	hold_number int
}

type student struct {
	info    Info
	contect Contect
	address Address
}

func main() {

	////////////
	//type 1///
	///////////
	var user User
	user.id = 1
	user.name = "Nahid"
	user.number = 15
	user.city = "Tangail"

	fmt.Println(user)

	///////////
	//type 2////
	//////////////
	user1 := User{
		id:     2,
		name:   "nahid",
		city:   "Tangail",
		number: 12345657}

	fmt.Println(user1)

	////////////
	//type 3///
	///////////
	var user3 = new(User)
	user3.id = 3
	user3.name = "Nahid"
	user3.number = 56415
	user3.city = "Tangail"

	fmt.Println(user3)

	//////////////////
	//complex strut//
	/////////////////
	var student student

	student.info = Info{
		name:  "Nahid",
		fName: "kader",
		mName: "Nasima",
	}
	student.contect = Contect{
		gmail:  "nahid@gmail.com",
		number: 234234,
	}
	student.address = Address{
		country:     "Bangladesh",
		city:        "Tangail",
		hold_number: 64,
	}

	fmt.Println(student)

}
