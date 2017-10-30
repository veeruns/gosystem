package gosystem

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func Mapfunction() {
	x := make(map[string]int)
	x["s1"] = 1
	x["s2"] = 2
	fmt.Println(x)

	if v, ok := x["s2"]; ok {
		fmt.Println(v)
	}

	delete(x, "s2")
	fmt.Println(x)

}

func Structfunction() {
	veeru := Person{
		Name:    "Veeru Natarajan",
		Age:     39,
		Address: "Fremont",
	}

	fmt.Println(veeru.Name)

}
