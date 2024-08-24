package main

import "fmt"

func main() {
	// string
	var nameOne string = "mario"
	var nameTwo string

	fmt.Println(nameOne, nameTwo)

	nameOne = "JoJo"
	nameTwo = "Josuke"

	fmt.Println(nameOne, nameTwo)

	nameThree := "giorno"

	fmt.Println(nameThree)

	var age1 int = 33
	var age2 = 19
	age3 := 40

	fmt.Println(age1, age2, age3)

	// number from -128 to 127
	var num1 int8 = 26
	var num2 int8 = -128

	// number which biger then 0
	var num3 uint = 33

	fmt.Println(num1, num2, num3)

	// #################################################################################################################################

	// arrays

	// biz tortburchakni ichidagi kiritilgan index sonini ozgartira olmaymiz

	numbers := [6]int{1, 2, 3, 4, 5, 16}
	numbers[2] = 2008

	// index chagarasiz

	var names = []string{"josuke", "giorno", "jolyne"}

	// add something
	names = append(names, "jotaro")

	fmt.Println(numbers, len(numbers))
	fmt.Println(names, len(names))

	// slices

	ranges1 := names[1:3]
	ranges2 := names[1:]
	ranges3 := names[:2]

	fmt.Println(ranges1, ranges2, ranges3)

	ranges1 = append(ranges1, "joseph")

	fmt.Println(ranges1)

	// #################################################################################################################################

	// prints

	fmt.Print("hello,", " ")
	fmt.Print("world!")

	fmt.Println("new line")

	// formatted print
	fmt.Printf("my name is %v my age is %v \n", names[0], numbers[5])

	// type ov varriables
	fmt.Printf("my name on type %T \n", names[0])
	fmt.Printf("my age on type %T \n", numbers[5])

	// print with symbols
	fmt.Printf("my name is %q my age is %q \n", names[0], numbers[5])

	// print floats

	fmt.Printf("your soursed %f points \n", 245.55)
	fmt.Printf("your soursed %0.1f points \n", 245.55)

	// print format on variable
	var str = fmt.Sprintf("my name is %v my age is %v \n", names[0], numbers[5])

	fmt.Println("saved string is:", str)

}
