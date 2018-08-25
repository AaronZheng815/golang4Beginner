package main

import (
	"fmt"
	"os"
	"time"
)

type PathError struct{
	path string
	op string
	ctime string
	msg string
}
func (p *PathError) Error() string{
	return fmt.Sprintf("\n%s\n%s\n%s\n%s\n", p.path, p.op, p.ctime,p.msg)
}

func OpenFile(fn string) error {
	
	file,err := os.Open(fn)
	defer file.Close()

	if err != nil {
		return &PathError{
			path : fn,
			op : "open",
			ctime : time.Now().String(),
			msg : err.Error(),
		}
	}

	return nil
}

func main(){
	err := OpenFile("test.sss")
	if err != nil {
		//fmt.Println("open filed with error: ", err)
		switch v:= err.(type){
		case *PathError:
			fmt.Println("get path error: ", v)
		default:
		}
	
		return 
	}


}