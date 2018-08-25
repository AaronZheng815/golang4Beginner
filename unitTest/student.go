package main

import (
	"io/ioutil"
	"encoding/json"
)

type Student struct{
	Name string
	Age int
	Sex string
}

func (p *Student) Save()(err error){
	data,err := json.Marshal(p)
	if err != nil{
		return
	}

	err = ioutil.WriteFile("stu.txt",data,0755)
	return
}

func (p *Student) Load()(err error){
	data,err := ioutil.ReadFile("stu.txt")
	if err != nil {
		return
	}

	err = json.Unmarshal(data,p)
	return
}
