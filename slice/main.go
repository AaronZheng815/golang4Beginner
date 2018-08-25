package main

import (
	"fmt"
)

func main()  {
	//切片定义
	var s1 []int
	arr := [...]int{0,1,2,3,4,5}
	
	//切片半开，包含2，不包含3
	s1 = arr[1:4]
	fmt.Println(arr)
	fmt.Println(s1)
	//cap表示的是从起始下标到数组最大下标的长度
	//len表示的是切片的长度
	fmt.Printf("s1 cap = %d, len =%d\n",cap(s1),len(s1))

	s1 = s1[0:2]
	fmt.Println(s1)
	fmt.Printf("s1 cap = %d, len =%d\n",cap(s1),len(s1))

	//简写
	s2 := arr[:]
    s3 := arr[:4]
	s4 := arr[2:]
	s5 := s2[:len(s2)-1]
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

	//通过make来创建切片
	//make([]type,len,cap)
	var s6 []int = make([]int, 5)
	s7 :=make([]int,5)
	s8 :=make([]int,6,8)
	fmt.Printf("s6: len = %d, cap = %d\n",len(s6),cap(s6))
	fmt.Printf("s7: len = %d, cap = %d\n",len(s7),cap(s7))
	fmt.Printf("s8: len = %d, cap = %d\n",len(s8),cap(s8))
	
	//创建切片
	strs := []string {"aaron","nina","allen","cici"}
	fmt.Println(strs)
	
	//append,会分配内存，扩充大小
	s6 = append(s6,1)
	s6 = append(s6,2)
	fmt.Println("s6: ", s6)
	fmt.Printf("s6: len = %d, cap = %d\n",len(s6),cap(s6))

	//查看append后地址的变化
	arr2 := [...]int{1,2,3,4,5}
	s9 := arr2[1:]
	fmt.Println("arr2: ",arr2,"arr2[1] addr: ",&arr2[1], " s9:",s9)
	fmt.Printf("ptr s9： %p\n",s9)
	//s9 = append(s9,10)
	s9 = s9[1:3]
	fmt.Println("arr2: ",arr2,"arr2[1] addr: ",&arr2[1], " s9:",s9)
	fmt.Printf("ptr s9： %p\n",s9)

	//copy
	var arr3 = [...]int{1,2,3}
	s10 :=arr3[:]
	s11 :=make([]int,10) 
	copy(s11,s10)         //如果长度不够，则只拷贝s11长度的元素
	fmt.Println(s10,s11)

	s12 :=[]int{5,5}
	s12 = append(s12,s10...) //通过...表示扩展
	fmt.Println(s12)
	s12 = append(s12,9,9,9)
	fmt.Println(s12)

	//字符串的改变
	str :="hello world"
	s13 :=[]byte(str)  //  []rune(str)  如果需要处理中文
	s13[0] = '0'
	str = string(s13)
	fmt.Println(str)

}