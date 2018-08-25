package main

import (
	"io"
	"bufio"
	"fmt"
	"os"
)

type CharCount struct{
	CharCount int
	SpaceCount int
	NumCount int
	OtherCount int
}

func countFile(fn string){
	file,err := os.Open(fn)
	if err != nil {
		fmt.Println("Faile to open file, error:", err)
		return
	}
	defer file.Close()

	var count CharCount
	reader := bufio.NewReader(file)
	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Printf("read file failed, err:%s\n",err)
			break
		}

		runeArr := []rune(str)
		for _,v := range runeArr{
			switch {
			case v >= 'a' && v <='z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.CharCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >='0' && v<='9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}

	//display the result
	fmt.Printf("Char count: %d\n",count.CharCount)
	fmt.Printf("Number count: %d\n",count.NumCount)
	fmt.Printf("Space count: %d\n",count.SpaceCount)
	fmt.Printf("Other count: %d\n",count.OtherCount)
}