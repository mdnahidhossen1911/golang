package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "nahid nafiz"
	namelist := strings.Split(name, " ")
	fmt.Println(namelist[0])
	fmt.Println(namelist)

	//start end space remove
	data := "   nahid hossen  "
	trimmed := strings.TrimSpace(data)
	fmt.Println(trimmed)

	//string join
	st1 := "nahid"
	st2 := "hossen"
	st3 := "abdullah"
	sjoin := strings.Join([]string{"MD", st1, st2, st3}, " ")
	fmt.Println(sjoin)
}
