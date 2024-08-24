package main

import "fmt"

func main() {
	// age := 55

	// fmt.Println(age == 33)
	// fmt.Println(age != 33)
	// fmt.Println(age > 45)
	// fmt.Println(age < 45)
	// fmt.Println(age >= 45)
	// fmt.Println(age <= 45)

	// if age == 33 {
	// 	fmt.Println("age is equal to 33")
	// } else if age < 42 {
	// 	fmt.Println("age is less then 42")
	// } else {
	// 	fmt.Println("idk")
	// }

	arr := []string{"jojo", "jonathan", "joseph", "jotaro", "josuke", "giorno", "jolyne"}

	for index, value := range arr {
		if index == 2 {
			fmt.Println("continuing on index", index)
			continue
		}

		if index == 4 {
			fmt.Println("breaking on index", index)
			break
		}

		fmt.Printf("the value at %v is %v \n", index, value)
	}
}
