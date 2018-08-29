package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	// c, err := redis.Dial("tcp", "localhost:6379")
	c, err := redis.Dial("tcp", "172.17.0.3:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	defer c.Close()

	//set int
	_, err = c.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	//get int
	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(r)

	//set string
	_, err = c.Do("Set", "stu_name", "Aaron Zheng")
	if err != nil {
		fmt.Println(err)
		return
	}
	//get string
	var str string
	str, err = redis.String(c.Do("Get", "stu_name"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println("get from reids, stu_name: ", str)
	
	//bach set
	_, err = c.Do("MSet", "aaa",100,"bbb","test")
	if err != nil{
		fmt.Println(err)
		return
	}
	rst, err := redis.Strings(c.Do("MGet","aaa","bbb"))
	if err != nil{
		fmt.Println(err)
		return
	}
	for _,v := range rst{
		fmt.Println("MGet: ", v)
	}

	//set hash
	strMap := make(map[string] interface{})
	strMap["name"] = "Aaron"
	strMap["age"] = 123
	fmt.Println(strMap)

	for k,v := range strMap{
		_,err = c.Do("HSet", "infoFeilds", k,v)
		if err != nil{
			fmt.Println(err)
			return
		}
	}
	//get hash
	str,err = redis.String(c.Do("HGet", "infoFeilds","name"))
	if err != nil{
		fmt.Println("get name failed, ", err)
		return 
	}
	fmt.Println("HGet, name: ", str)
	
	age,err := redis.Int(c.Do("HGet", "infoFeilds","age"))
	if err != nil{
		fmt.Println("get name failed, ", err)
		return 
	}
	fmt.Println("HGet age: ", age)

	//list
	_, err = c.Do("lpush", "list01", "abc", "def", 300, 400)
	if err != nil {
		fmt.Println(err)
		return
	}

	len, _ := redis.Int(c.Do("LLen","list01"))
	for i:=0;i<len;i++{
		str, err = redis.String(c.Do("lpop", "list01"))
		if err != nil {
			fmt.Println("get list failed,", err)
			return
		}
		fmt.Println(str)
	}

	//expire
	_, err = c.Do("expire", "list01", 10)
	if err != nil {
		fmt.Println(err)
		return
	}


}
