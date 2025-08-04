package tut
import (
	"fmt"
)

func inc (n *int) {
	*n++
}

func One () {
	x := 5
	inc (&x)
	fmt.Println (*(&x))
}
