package main

import (
	"fmt"
)

func facotrial(num uint64) uint64  {
	defer func(){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()
	
	if num == 1 {
		return num
	}

	return facotrial(num - 1) * num
}

func fibonacci(num int) int {
	if num <= 1 {
		return 1
	}
	if num == 2 {
		return 2
	}

	return fibonacci(num - 1) + fibonacci(num - 2)
}

func main()  {
	var num int;
	fmt.Print("Please input a number less then 65: ")
	fmt.Scanf("%d",&num)

	fmt.Printf("The factorial of %d is %d.\n",
				num,
				facotrial(uint64(num)))

	fmt.Printf("The fibonacci number of %d is %d.\n",
				num,
				fibonacci(num))
}