package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}

func addNode(headPtr,stu *Student){
	// tmp := headPtr
	// for tmp.next != nil {
	// 	tmp = tmp.next
	// }
	// tmp.next = stu

	for headPtr.next != nil{
		headPtr = headPtr.next
	}

	headPtr.next = stu
}

func displayList(nodePtr *Student){
	// tmp := head

	// for tmp.next != nil {
	// 	fmt.Println(tmp.Name,tmp.Age,tmp.Score)
	// 	tmp = tmp.next
	// }
	// fmt.Println(tmp.Name,tmp.Age,tmp.Score)	

	// for{
	// 	fmt.Println(tmp.Name,tmp.Age,tmp.Score)

	// 	if tmp.next != nil{
	// 		tmp = tmp.next
	// 	}else{
	// 		break
	// 	}
	// }

	for nodePtr != nil {
		fmt.Println(nodePtr.Name,nodePtr.Age,nodePtr.Score,nodePtr.next)
		nodePtr = nodePtr.next
	}
}

func main() {
	var head Student
	head.Name = "xiaoming"
	head.Age = 10
	head.Score = 8.5

	var stu1 Student
	stu1.Name = "xiaohe"
	stu1.Age = 11
	stu1.Score = 6.7

	head.next = &stu1

	var stu2 Student
	stu2.Name = "xiaogao"
	stu2.Age = 12
	stu2.Score = 8.7

	var stu3 Student
	stu3.Name = "xiaochen"
	stu3.Age = 13
	stu3.Score = 6.7

	addNode(&head,&stu2)
	addNode(&head,&stu3)
	displayList(&head)
}
