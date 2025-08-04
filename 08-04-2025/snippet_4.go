package tut
import (
	"fmt"
)

type PersonFour struct {
	Name string
	Age int
}

func Four () {
	var p1 PersonFour = PersonFour {"Rahul", 21}
	fmt.Println (p1)

	p2 := PersonFour {Name: "Navneeth", Age: 21}
	fmt.Println (p2)

	p2.Age = 22
	fmt.Println (p2.Name, p2.Age)
}
