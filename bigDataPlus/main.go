package main

import (
	"strings"
	"fmt"
	"os"
	"bufio"
)

func addBig(str1,str2 string) (result string){
	if len(str1) == 0 && len(str2) == 0{
		result = "0"
		return
	}

	var index1 = len(str1) - 1
	var index2 = len(str2) - 1
    var left int
	for index1 >=0 && index2 >=0 {
		c1 := str1[index1] - '0'
		c2 := str2[index2] - '0'
		sum := int(c1) + int(c2) + left

		if (sum >= 10){
			left = 1
		} else{
			left = 0
		}

		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s",c3,result)

		index1 --
		index2 --
	} 

	for index1 >= 0 {
		c1 := str1[index1] - '0'
		sum := int(c1) + left
		
		if (sum >= 10){
			left = 1
		} else{
			left = 0
		}

		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s",c3,result)

		index1 --
	}

	for index2 >= 0 {
		c2 := str2[index2] - '0'
		sum := int(c2) + left
		
		if (sum >= 10){
			left = 1
		} else{
			left = 0
		}

		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s",c3,result)

		index2 --
	}
	return
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from concole err: ", err)
		return
	}

	strSlice := strings.Split(string(result), "+")
	if(len(strSlice) != 2){
		fmt.Println("Please input a+b")
		return
	}
	strNum1 := strings.TrimSpace(strSlice[0])
	strNum2 := strings.TrimSpace(strSlice[1])
	
	fmt.Println("result is = ", 
	             addBig(strNum1,strNum2))
}