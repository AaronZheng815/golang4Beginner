package main

import (
	"testing"
)



func TestSave(t *testing.T){
	stu :=&Student{
		Name:"Aaron",
		Age:18,
		Sex:"male",
	}

	err := stu.Save()
	if err != nil{
		t.Fatalf("Save student failed.")
	}

}

func TestLoad(t *testing.T){
	stu :=&Student{
		Name:"Aaron",
		Age:18,
		Sex:"male",
	}
	err := stu.Save()
	if err != nil{
		t.Fatalf("Save student failed.")
	}

	stu2 := &Student{}
	err = stu2.Load()
	if err != nil{
		t.Fatalf("load Student Failed, err:%v",err)
	}

	if stu.Name != stu2.Name{
		t.Fatalf("load Student Failed, name dismatch")
	}
	
}