package loadbalance

import (
	"fmt"
)

type LbManager struct{
	allLBAlgo map[string] LoadBalance
}

var (
	lbMgr = LbManager{
		allLBAlgo : make(map[string]LoadBalance),
	}
)

func (p *LbManager) registerLbAlgo(name string, lb LoadBalance){
	p.allLBAlgo[name] = lb
}

func RegisterLbAlgo(name string, lb LoadBalance){
	lbMgr.registerLbAlgo(name, lb)
}

func DoLB(name string, insts[]*Instance, key ...string) (inst *Instance, err error){
	balancer,ok := lbMgr.allLBAlgo[name]
	if !ok {
		err = fmt.Errorf("Not found.%s",name)
		return
	}

	if len(key) == 0{
		inst, err = balancer.DoBalance(insts)
	} else {
		inst, err = balancer.DoBalance(insts, key[0])
	}
	
	return
}