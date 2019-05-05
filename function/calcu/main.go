package main

import "fmt"

type operator func(x, y int) int

func calculate(x, y int, op operator) (result int, err error) {
	if op == nil {
		return 0, fmt.Errorf("invalid operator")
	}

	result = op(x, y)

	return
}

func main() {

	x, y := 1000, 2000

	add := func(x, y int) int {
		return x + y
	}

	rt, err := calculate(x, y, add)
	fmt.Printf("%d + %d = %d, error(%v)\n", x, y, rt, err)

	multiply := func(x,y int)int{
		return x * y
	}

	rt, err = calculate(x, y, multiply)
	fmt.Printf("%d x %d = %d, error(%v)\n", x, y, rt, err)
}
