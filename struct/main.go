package main

import (
	"time"
	"fmt"
)
//定义
type Student struct{
	name string
	age int
}

//别名
type Stu Student

//匿名字段, golong中通过它来实现继承
type Dog struct{
	name string
	age int
}
type MyDog struct{
	Dog
	int
	start time.Time
}

func NewStudent(name_r string, age_r int) *Student  {
	return &Student{
		name:name_r,
		age:age_r,
	}
}

//结构体方法
// func (reciver type) methodName (参数列表) (返回值列表)
// 注意reciver的值传递和指针传递
// 方法的访问权限也是通过大小写来控制

func (p *Student) init(name string, age int){
	p.name = name
	p.age = age
	fmt.Println(p)
}
func (self *Student) getName() string{
	return self.name
}
func (self *Student) setName(name string) {
	self.name = name
}

//普通类型的方法
type integer int
func (p integer) print(){  //接受者这里，也是类似函数参数传递，需要区分值传递和指针传递
	fmt.Println(p)
}
func (p *integer) set(num integer){
	*p = num
}

func main()  {

	//定义变量
	var stu Student
	stu.name = "Aaron Zheng"   //初始化
	stu.age = 18

	fmt.Println(stu)
	
	var stu1 *Student = new(Student)
	var stu2 *Student = &Student{"bbb",2}  //初始化

	stu1.name = "aaa"  
	stu1.age = 1
	fmt.Println(stu1.age,stu2.age)

	stu3 := Student{name:"nina",age:18}    //初始化,可以只初始化部分字段
	fmt.Println(stu3)

	stu4 := Stu{name:"cici"}
	fmt.Println(stu4)

	//golong中没有构造函数，可以使用工厂模式来创建
	stu5 := NewStudent("allen",5)
	fmt.Println(stu5)

	//访问匿名字段,有些类似于C++中基类的方式
	var x MyDog
	x.name = "xiaobai"
	x.age = 1
	x.int =100
	x.start = time.Now()

	fmt.Println(x)

	//使用方法
	var stu6 Student
	stu6.init("xiaoxiao",18) // == (&stu6).init("xiaoxiao",18)
	//(&stu6).init("xiaoxiao",18)
	fmt.Println(stu6)
	fmt.Println(stu6.getName())
	stu6.setName("tiantian")
	fmt.Println(stu6.getName())

	//普通自定义类型的方法
	var inta integer
	inta = 100
	inta.print()
	inta.set(200)
	inta.print()
}
