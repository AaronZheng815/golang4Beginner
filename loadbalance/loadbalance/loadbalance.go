package loadbalance

type LoadBalance interface{
	DoBalance([]*Instance, ...string) (inst *Instance, err error)
}