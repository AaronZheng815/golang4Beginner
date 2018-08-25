package main

import (
	"os"
	"text/template"
	"fmt"
)

type Person struct{
	Name string
	Age int
}

func main()  {
	t, err := template.ParseFiles("person.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}

	p := Person{ Name: "aaron", Age:18}

	if err:=t.Execute(os.Stdout,p); err != nil{
		fmt.Println("There are errors:",err)
	}

	
}