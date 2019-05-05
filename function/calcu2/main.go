package main

import (
	"errors"
	"fmt"
)

type operator func(x, y int) int

type calculateFunc func(x, y int) (result int, err error)

func genCalculator(op operator) calculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {

	x, y := 1000, 2000

	op := func(x, y int) int {
		return x + y
	}

	add := genCalculator(op)
	rt, err := add(x, y)

	fmt.Printf("%d + %d = %d, error(%v)\n", x, y, rt, err)

}
