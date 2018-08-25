package main

import (
	"fmt"
	"encoding/json"
)

type User struct{
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age int
	Birthday string
	Sex string
	Email string
	PhoneNum string
}

func structPack(){
	user :=&User{
		UserName : "XiaoMin",
		NickName : "Aaron",
		Age : 18,
		Birthday : "2018-08-26",
		Sex : "male",
		Email : "qq@qq.com",
		PhoneNum : "110",
	}

	data,err := json.Marshal(user)
	if err != nil{
		fmt.Println("json marshal failed: ", err)
		return
	}
	fmt.Println("Json: ", string(data))

	encodeStr := string(data)
	var user1 User
	err = json.Unmarshal([]byte(encodeStr), &user1)
	if err != nil {
		fmt.Println("json unmarshal failed: ",err)
	}

	fmt.Println("struct: ", user1)

}

func intPack(){
	i := 110
	data,err := json.Marshal(i)
	if err != nil{
		fmt.Println("json marshal failed: ", err)
		return
	}
	fmt.Println("json: ", string(data))

	encodeStr := string(data)
	var user1 int
	err = json.Unmarshal([]byte(encodeStr), &user1)
	if err != nil {
		fmt.Println("json unmarshal failed: ",err)
	}

	fmt.Println("int: ", user1)
}

func mapPack(){
	var m map[string]interface{}

	m = make(map[string]interface{})
	m["uname"] = "xiagnwang"
	m["age"] = 18
	m["flag"] = true
	
	data,err := json.Marshal(m)
	if err != nil{
		fmt.Println("json marshal failed: ", err)
		return
	}
	fmt.Println("json: ", string(data))

	encodeStr := string(data)
	var user1 map[string]interface{}
	err = json.Unmarshal([]byte(encodeStr), &user1)
	if err != nil {
		fmt.Println("json unmarshal failed: ",err)
	}

	fmt.Println("map: ", user1)
}

func slicePack(){
	var s []map[string]interface{}
	var m map[string]interface{}

	m = make(map[string]interface{})
	m["uname"] = "xiagnwang"
	m["age"] = 18
	m["flag"] = true
	s = append(s,m)

	m = make(map[string]interface{})
	m["uname"] = "xiao"
	m["age"] = 20
	m["flag"] = false
	s = append(s,m)

	data,err := json.Marshal(s)
	if err != nil{
		fmt.Println("json marshal failed: ", err)
		return
	}
	fmt.Println("json: ", string(data))

	encodeStr := string(data)
	var user1 []map[string]interface{}
	err = json.Unmarshal([]byte(encodeStr), &user1)
	if err != nil {
		fmt.Println("json unmarshal failed: ",err)
	}

	fmt.Println("slice: ", user1)
}

func main()  {
	structPack()

	intPack()
	
	mapPack()
	
	slicePack()
}