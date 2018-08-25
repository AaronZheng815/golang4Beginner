package main

import (
	"flag"
	"os"
	"fmt"
)

func main()  {
	fmt.Println(os.Args)

	lens := len(os.Args)
	fmt.Printf("there are %d Para.\n", lens)
	for index, v := range os.Args{
		fmt.Printf("args[%d] = %v", index, v)
	}

	//命令行使用
	fmt.Println()
	var confPath string
	flag.StringVar(&confPath, "c", " ", "Please input config file")
	
	var logLevel int
	flag.IntVar(&logLevel, "l", 1, "Please input log level(0-3)")
	
	flag.Parse()
	fmt.Println("conf file: ",confPath," log level: ", logLevel)

	return
}