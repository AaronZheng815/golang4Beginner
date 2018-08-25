package main

import (
	"fmt"
)

//如果一个数恰好等于它的因子之和，则称该数为“完全数”。
//第一个完全数是6，它有约数1、2、3、6，除去它本身6外，其余3个数相加，1+2+3=6。
func isPerfectNum(n int) bool{
	var sum int
	for i:=1;i<n;i++{
		if(n%i==0){
			sum += i
		}
	}
	return n == sum
}

func perfectNum(n int)  {
	for i:=1;i<=n;i++{
		if isPerfectNum(i) == true{
			fmt.Println(i)
		}
	}
}

func main()  {
	var n int
	fmt.Scanf("%d", &n)
	perfectNum(n)
}