package main

import "fmt"

func main() {
	x := 0

	for x <= 5 {
		fmt.Println("loop -", x)
		x++
	}

	for i := 0; i <= 100; i++ {
		fmt.Println("loop number -", i)
	}

	arr := []string{"jojo", "jonathan", "joseph", "jotaro", "josuke", "giorno", "Jolyne"}

	for i := 0; i < len(arr); i++ {
		fmt.Println("on index", i, "-", arr[i])
	}

	// with all properties
	for index, value := range arr {
		fmt.Printf("the value on index %v is %v \n", index, value)
	}

	// removing property which we dont use
	for _, value := range arr {
		fmt.Printf("the value is %v \n", value)
	}

}
