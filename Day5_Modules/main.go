package main

import (
	"fmt"

	sliceHelper "github.com/ankitkr2419/sliceHelper"
)

func execute() {
	s1 := []string{"a", "b", "c", "d"}
	s2 := []string{"a", "b", "g", "k"}

	ele := "b"

	isPresent := sliceHelper.FindElement(ele, s1)
	diff := sliceHelper.DiffElements(s1, s2)
	unique := sliceHelper.UniqueElements(s1)
	fmt.Println(isPresent)
	fmt.Println(diff)
	fmt.Println(unique)
}

func main() {
	execute()
}
