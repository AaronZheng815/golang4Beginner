package main

import (
	"time"
	"fmt"
	"net/http"
)

func rootProcess(w http.ResponseWriter, r *http.Request){
	fmt.Println("handle root request")
	fmt.Fprintf(w,"hello go http")
}

func loginHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("handle login request")
	fmt.Fprintf(w,"login go http")
}

func displayHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("handle login request")
	fmt.Fprintf(w,"display")
}

// func httpSer(){
// 	http.HandleFunc("/", rootProcess)
// 	http.HandleFunc("/login", loginHandler)
// 	http.HandleFunc("/display",displayHandler)
// 	err := http.ListenAndServe("0.0.0.0:5000",nil)
// 	if err != nil {
// 		fmt.Println("http listen failed")
// 	}
// }
func main()  {
	// go httpSer()

	// for{
	// 	time.Sleep(time.Second)
	// 	fmt.Println("main process, echo")
	// }

	http.HandleFunc("/", rootProcess)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/display",displayHandler)
	err := http.ListenAndServe("0.0.0.0:5000",nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}