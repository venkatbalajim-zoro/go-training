package tut
import (
	"fmt"
)

func update (n *int) {
	*n = 2
}

func Two () {
	x := 1
	fmt.Println (x)
	update (&x)
	fmt.Println (x)
}
