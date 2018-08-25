package main

import (
	"sort"
	"math/rand"
	"fmt"
)

type Student struct{
	Name string
	Id string
	Age int
}

type Book struct{
	Name string
	Author string
}

type StudentArray []Student

func (p StudentArray) Len() int{
	return len(p)
}

func (p StudentArray)  Less(i,j int) bool {
	return p[i].Name < p[j].Name
}

func (p StudentArray) Swap(i,j int){
	p[i], p[j] = p[j], p[i]
}

func main()  {
	var stus StudentArray
	for i:=0;i<10; i++{
		stu := Student{
			Name : fmt.Sprintf("student%d", rand.Intn(100)),
			Id : fmt.Sprintf("1000%d",rand.Intn(100)),
			Age : rand.Intn(100),
		}
		stus = append(stus, stu)
	}

	for _,v := range stus{
		fmt.Println(v)
	}
	
	sort.Sort(stus)
    fmt.Println("After sort:")
	for _,v := range stus{
		fmt.Println(v)
	}

}