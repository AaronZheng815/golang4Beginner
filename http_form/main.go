package main

import (
	"text/template"
	"log"
	"io"
	"fmt"
	"net/http"
)

type Person struct{
	Name string
	Age int
}

const form = `<html><body><form action="#" method="post" name="bar">
<input type="text" name="usrname"/>
<input type="text" name="passwd"/>
<input type="submit" value="Submit"/>
</form></body></html>
`

func main()  {
	http.HandleFunc("/hw", logPanics(hwHandler))
	http.HandleFunc("/login", logPanics(loginHandler))
	err := http.ListenAndServe(":8888", nil)
	if err != nil{
		fmt.Println("http server corrupt")
	}
	return
}

func hwHandler(w http.ResponseWriter, req *http.Request)  {
	hwStr := "<h1>Hello World</h>"
	io.WriteString(w,hwStr)
	panic("error panic")
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch req.Method{
	case "GET":
		io.WriteString(w,form)
	case "POST":
		fmt.Printf("username = %s, passwd = %s\n",
				req.Form["usrname"][0],
				req.FormValue("passwd"))
		// io.WriteString(w,"<h1>Sucess!</h1>")
		username := req.FormValue("usrname")
		
		t, err := template.ParseFiles("person.html")
		if err != nil {
			fmt.Println("parse file err:", err)
			return
		}
	
		p := Person{ Name: username, Age:18}

		if err:=t.Execute(w,p); err != nil{
			fmt.Println("There are errors:",err)
		}
	
	}
}

func logPanics(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}
