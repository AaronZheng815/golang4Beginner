package main

import (
	"fmt"
)

type Reader interface{
	Read()
}

type Writer interface{
	Write()
}

type ReadWriter interface{
	Reader
	Writer
}

type File struct{

}

func (f *File) Read(){
	fmt.Println("read Data")
}

func (f *File) Write(){
	fmt.Println("write Data")
}

func Test( rw ReadWriter){
	rw.Read()
	rw.Write()
}

func main()  {

	var f File
	Test(&f)		
}