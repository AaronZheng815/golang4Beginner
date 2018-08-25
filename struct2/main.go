package main

import (
	"fmt"
)

type Car struct{
	weight int
	colar string
}
func (this *Car) Run(){
	fmt.Println("Car Running")
}

type Bike struct{
	Car
	wheel int
}

type Train struct{
	Car
	wheel int
}

//组合
type Transport struct{
	a Bike
	b Train
}

func main(){
	var a Bike
	a.weight = 100
	a.colar = "red"
	a.wheel = 2
	fmt.Println(a)
	a.Run()

	var b Train
	b.weight = 200
	b.colar = "green"
	b.wheel = 1111
    b.Run()
}