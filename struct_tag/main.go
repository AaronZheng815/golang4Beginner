package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	// Name string
	// Age int
	// Score int
	Name  string `json:"student_name"`
	Age   int    `json:"student_age"`
	Score int    `json:"student_score"`
}

func main() {
	stu := Student{"aaron", 18, 90}

	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("Json encode failed.err: ", err)
	}

	fmt.Println(string(data))

}
