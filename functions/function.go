package main

import (
	"fmt"
	"strings"
)

func seyHi(name string) {
	fmt.Printf("ciao, %v \n", name)
}
func seyBye(name string) {
	fmt.Printf("arrividarcie, %v \n", name)
}

func loopFunctions(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func CL(t string) (string, string) {
	CLT := strings.ToUpper(t)
	splited := strings.Split(CLT, " ")

	var inits []string
	for _, v := range splited {
		inits = append(inits, v[:1])
	}

	if len(inits) > 1 {
		return inits[0], inits[1]
	}

	return inits[0], "_"
}

func main() {
	// seyHi("Mardonbek")
	// seyBye("Mardonbek")

	// loopFunctions([]string{"josuke", "giorno", "jolyne"}, seyBye)

	capL := []string{}

	fm1, fn1 := CL("mardonbek khamidov")

	capL = append(capL, fm1, fn1)

	fm2, fn2 := CL("mardonbek")

	fmt.Println(capL)
	println(fm2, fn2)

	//  := CL("mardonbek khamidov")

}
