package main

import (
	"compress/gzip"
	"io/ioutil"
	"bufio"
	"os"
	"fmt"
)

type Student struct{
	Name string
	Age int
	Score float32
}

func main()  {
    //read from string
	var str = "stu01 18 90.5"
	var stu Student
	fmt.Sscanf(str,"%s %d %f", &stu.Name, &stu.Age, &stu.Score)
	fmt.Println(stu)

	//read from stdin
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read failed: ",err)
		return 
	}
	fmt.Printf("read str, ret: %s\n",str)
	
	//read from file
	file,err := os.Open("test.txt")
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}
	defer file.Close()
	reader1 := bufio.NewReader(file)
	str1, err := reader1.ReadString('\n')
	if err != nil {
		fmt.Println("read failed: ",err)
		return 
	}
	fmt.Printf("read str, ret: %s\n",str1)

	countFile("test.txt")

	//read with iotuil
	inFile := "test.txt"
	outFile := "test_bak.txt"

	buf,err:=ioutil.ReadFile(inFile)
	if err != nil{
		fmt.Println("read file failed: ",err)
		return
	}

	fmt.Printf("%s\n",buf)
	err = ioutil.WriteFile(outFile, buf, 0x64)
	if err != nil{
		panic(err.Error())
	}

	//读取压缩文件
	fmt.Println("-----------start read gz file--------------")
	fName := "test.gz"
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, can't open %s: error: %s\n",os.Args[0],fName,err)
		os.Exit(1)
	}

	fz,err := gzip.NewReader(fi)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open gzip failed, err: %v\n", err)
		return
	}

	r := bufio.NewReader(fz)
	for{
		line,err:=r.ReadString('\n')
		if err != nil{
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}