package main

import (
	"fmt"
)

type Test interface{
	Print()
	Sleep()
}

type Student struct{
	name string
	age int
	score int
}

func (this *Student) Print(){
	fmt.Println(this.name)
	fmt.Println(this.age)
	fmt.Println(this.score)
}

func (this *Student) Sleep(){
	fmt.Println(this.name, " is sleeping")
}

type People struct{
	name string
	age int
}
func (pep *People) Print(){
	fmt.Println(pep.name)
	fmt.Println(pep.age)
}

func (pep People) Sleep(){
	fmt.Println(pep.name, " is sleeping")
}

func main()  {
	var t Test
	var stu Student = Student{"xiaoming", 12,98}
	
	t = &stu
	t.Print()
	t.Sleep()

	var pep = People{"people",12}
	t = &pep
	t.Print()
	t.Sleep()
}