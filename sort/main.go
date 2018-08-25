package main

import (
	"fmt"
	"sort"
)

func intSortTest()  {
	var a = [...]int{1,66,2345,7,8,99,34,67}
	
	//slice := a[:]
	//sort.Ints(slice)	
	// fmt.Println(slice)

	sort.Ints(a[:]) //入参需要时切片
	fmt.Println(a)
}

func stringSortTest()  {
	//strs := make([]string,5)
    strs := []string {"aaron","nina","allen","cici"}
	fmt.Println(strs)
	sort.Strings(strs)
	fmt.Println(strs)
}

func floatSortTest(){
	floats := []float64{0.123,234,55.44,9.888}
	fmt.Println(floats)
	sort.Float64s(floats)
	fmt.Println(floats)

}

func intSearchTest()  {
	ints := []int{1,66,2345,7,8,99,34,67}

	sort.Ints(ints)
	index := sort.SearchInts(ints,99)
	fmt.Println(ints)
	fmt.Println("found with :", index)
}

func stringSearchTest(){
	strs := []string{"aaron","nina","cici","allen"}
	sort.Strings(strs)
	fmt.Println(strs)

	index := sort.SearchStrings(strs,"cici")
	fmt.Println("found with index: ",index)
}

func floatSearchTest()  {
	floats := []float64{0.123,234,55.44,9.888}
	sort.Float64s(floats)
	fmt.Println(floats)

	index := sort.SearchFloat64s(floats,55.44)
	fmt.Println("found with index: ",index)
}

func main()  {
	intSortTest()
	stringSortTest()
	floatSortTest()
	
	intSearchTest()
	stringSearchTest()
	floatSearchTest()
}