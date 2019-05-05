package main

import "fmt"

type Printer func(contents string) (n int, err error)
//type Printer func(string) (int, error)

func printToStd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}

func main() {
	var p Printer
	p = printToStd
	p("something")
}
