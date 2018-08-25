package loadbalance

import (
	"fmt"
	"errors"
)

func init(){
	fmt.Println("rondrobin algo has registerd.")
	RegisterLbAlgo("roundrobin",&RoundRobinBalance{})
}

type RoundRobinBalance struct {
	curIndex int
}

func (sp *RoundRobinBalance)DoBalance(insts []*Instance, key ...string) (inst *Instance, err error){
	if len(insts) == 0 {
		err = errors.New("No Instance")
		return 
	}

	lens := len(insts)
	if sp.curIndex >= lens {
		sp.curIndex = 0
	}

	index := sp.curIndex
	inst = insts[index]

	sp.curIndex = (sp.curIndex + 1) % lens
	return
}