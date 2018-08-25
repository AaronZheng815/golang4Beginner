package main

import (
	"fmt"
)

//传数组值
func handleArray(arr [5]int)  {
	arr[1] = 100
}

//传数组指针
func handleArray1(arr *[5]int){
	(*arr)[1] = 100
}

//打印斐波拉器数列
func fab(num int)  {
	var arr[] int
	arr = make([]int,num)

	arr[0] = 1
	arr[1] = 2

	for i:=2; i < num; i++{
		arr[i] = arr[i-1] + arr[i-2]
	}

	for _,v := range arr{
		fmt.Printf("%d ",v)
	}
}

func main()  {
	//数组的定义
	var arr1 [5]int
	arr1[0] = 1
	fmt.Println(arr1)

	//初始化
	arr2 := [5]int{1,2,3,4,5}
	fmt.Println(arr2)

	arr3 := [5]int{1,2}
	fmt.Println(arr3)

	arr4 := [...]int{1,2,3,4,5,6}
	fmt.Println(arr4)

	arr5 :=[5]string{1:"aaron", 4:"allen"}
	fmt.Println(arr5)

	//数组的下标遍历
	for i:=0;i<len(arr1);i++{
		fmt.Printf("%d ", arr1[i])
	}
	//数组的rang遍历
	for index,value := range arr1{
		fmt.Printf("arr1[%d] = %d ", index, value)
	}
    fmt.Println()
	
	//数组值传递，不改变内容
	handleArray(arr1)
	fmt.Println(arr1)

	handleArray1(&arr1)
	fmt.Println(arr1)

	//打印斐波拉器数列
	fab(50)
	fmt.Println()
	
	//多维数组
	var age [5][3]int
	fmt.Println(age)

	age1 := [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(age1)
	
	for row, vrow := range age1{
		for col, vcol := range vrow{
			fmt.Printf("a[%d][%d]=%d ",row,col,vcol)
		} 
		fmt.Println()
	}

}