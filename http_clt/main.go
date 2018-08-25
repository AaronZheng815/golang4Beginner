package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	res, err := http.Get("http://127.0.0.1:5000/display")
	if err != nil {
		fmt.Println("Get err: ", err)
		return
	}

	data, err :=ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Get data err:", err)
		return 
	}

	fmt.Println(string(data))

}