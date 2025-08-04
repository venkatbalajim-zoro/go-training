package tut
import (
	"fmt"
)

func Three () {
	x := 42
	var p *int = &x
	var pp **int = &p

	fmt.Println ("x:", x)
	fmt.Println ("*p:", *p)
	fmt.Println ("**p:", **pp)

	**pp = 100
	fmt.Println ("Updated x:", x)
}
