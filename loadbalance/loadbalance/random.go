package loadbalance

import (
	"fmt"
	"math/rand"
	"errors"
)

func init(){
	fmt.Println("rand algo has registerd.")
	RegisterLbAlgo("random",&RandomBalance{}) //相当于new了一个RandomBalance
}

type RandomBalance struct {

}

func (sp *RandomBalance)DoBalance(insts []*Instance, key ...string) (inst *Instance, err error){
	if len(insts) == 0 {
		err = errors.New("No Instance")
		return 
	}

	lens := len(insts)
	index := rand.Intn(lens)

	inst = insts[index]

	return
}