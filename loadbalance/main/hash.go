package main

import (
	"hash/adler32"
	"errors"
	"fmt"
	"basic/loadbalance/loadbalance"
)
type HashLoadBalance struct{
	key string
}

func init(){
	fmt.Println("hash alog has registered!")
	loadbalance.RegisterLbAlgo("hash", &HashLoadBalance{}) //相当于new了一个RandomBalance
}

func (sp *HashLoadBalance) DoBalance(insts []*loadbalance.Instance, key ...string) (inst *loadbalance.Instance, err error){
	if len(key) == 0{
		err = fmt.Errorf("hash param should be passed.")
		return
	}

	if len(insts) == 0 {
		err = errors.New("No Instance")
		return 
	}

	lens := len(insts)
	hashVal := adler32.Checksum([]byte(key[0]))
	index := int(hashVal) % lens

	inst = insts[index]

	return
}