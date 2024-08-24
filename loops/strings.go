package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("inital log __?")

	greeting := "String for experiment"

	// splitting and slising to arrays
	splited := strings.Split(greeting, " ")
	fmt.Println(splited)

	// checking words from str
	found := strings.Contains(greeting, "for")
	fmt.Println(found)

	// replasing all word
	replased := strings.ReplaceAll(greeting, "experiment", "exprerians")
	fmt.Println(replased)

	// uppercase
	up := strings.ToUpper(greeting)
	// lowercase
	low := strings.ToLower(greeting)
	// capitalise

	fmt.Println(up)
	fmt.Println(low)

	// find strign by index
	indexStr := strings.Index(greeting, "ex")
	fmt.Println(indexStr)

	// original

	// sort ###############################################################################################################################

	arr := []string{"jojo", "jonathan", "joseph", "jotaro", "josuke", "giorno", "Jolyne"}
	sort.Strings(arr)

	ints := []int{12, 23, 23, 34, 45, 545, 56, 5, 6}
	sort.Ints(ints)

	// finding a word from string
	index := sort.SearchStrings(arr, "Jolyne")

	fmt.Println(index)

	fmt.Println(arr)
	fmt.Println(ints)

	fmt.Println("original string =", greeting)

}
