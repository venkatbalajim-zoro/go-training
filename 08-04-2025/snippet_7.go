package tut
import (
	"fmt"
)

type Vehicle interface {
	Drive () string
	Rest () string
}

type Car struct {
	Model string
}

func (c Car) Drive () string {
	return "Carrrr!"
}

func (c Car) Rest () string {
	return "Zzzzzz!"
}

func checkVehicle (v Vehicle) {
	if car, ok := v.(Car); ok {
		fmt.Printf ("Car model : %s\n", car.Model)
	} else {
		fmt.Println ("Unknown type")
	}
}

func Seven () {
	v := Car {Model: "Audi"}
	fmt.Println (v)
	fmt.Println (v.Drive ())
	fmt.Println (v.Rest ())
	checkVehicle (v)
}
