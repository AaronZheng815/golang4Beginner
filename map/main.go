package main

import (
	"sort"
	"fmt"
)

func main()  {
	//key-value
	//var a map[int] string
	//声明的时候初始化
	var map1 = map[string]string{
		"aaa":"111",
		"bbb":"222",
	}
	fmt.Println(map1)

	//make初始化
	map2 := make(map[string] string)
	map2["aaron"] = "aaaa"
	map2["nina"] = "bbbb"
	map2["cici"] = "cccc"
	map2["allen"] = "dddd"
	fmt.Println(map2)
	
	map3 :=map[string] map[int]string{
		"aaa":{1:"aaron",2:"nina",},
		"bbb":{1:"cici",2:"allen",},
	}

	fmt.Println(map3)
	fmt.Println(map3["bbb"])

	//更新
	fmt.Println(map1)
	val,ok :=map1["ccc"]
	if ok {
		fmt.Println(val)
	} else {
		map1["ccc"] = "test"
		fmt.Println(map1)
	}

	//遍历
	for key,value := range map2{
		fmt.Println(key," = ", value)
	}

	//删除
	map2["xiaobai"] = "a doog"
	fmt.Println(map2)
	delete(map2,"xiaobai")
	fmt.Println(map2)
	fmt.Println(len(map2))
	//删除所有元素，需要循环

	//slice of map
	item := make([]map[string]string,5)
	for i:=0; i<5;i++ {
		item[i] = make(map[string]string)
	}

	item[0]["aaron"] = "aaaa"
	item[0]["nina"] = "bbbb"
	item[1]["aaron"] = "aaaa"
	fmt.Println(item)

	//map里面的key是无序的
	testMapSort()
}

func testMapSort(){
	a := make(map[int]int)
	a[2]=12
	a[1]=12
	a[3]=45
	a[66]=66

	//无序
	for k,v :=range a{
		fmt.Println(k,v)
	}

	//有序遍历
	fmt.Println("有序遍历")
	var keys []int
	for k,_ := range a{
		keys = append(keys,k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	for _,key :=range keys{
		fmt.Println(key,a[key])
	}
}

