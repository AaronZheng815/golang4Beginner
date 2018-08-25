package main

import (
	"math/rand"
	"fmt"
)

type Student struct{
	Name string
	Age int
}
func (p *Student) String() string{
	return fmt.Sprintf("name - %s, age = %d", p.Name,p.Age)
}

func main()  {
	var testlink Link

	for i :=0; i <10; i++{
		testlink.AddNodeBegin(i)
	}

	testlink.DisplayLink()

	var stringLink Link
	for i:=0; i<10; i++{
		str := fmt.Sprintf("str%d",rand.Intn(100))
		stringLink.AddNodeEnd(str)
	}
	stringLink.DisplayLink()

	var structLink Link
	for i:=0;i<10;i++{
		stu := &Student{
			Name:fmt.Sprintf("Student%d",rand.Intn(100)),
			Age:rand.Intn(30),
		}
		structLink.AddNodeBegin(stu)
	}
	structLink.DisplayLink()
}