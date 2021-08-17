package balance

import (
	"fmt"
)

//RoundRobin负载均衡
type RoundRobinBalance struct {
	curIndex int
}

func init() {
	RegisterBalancer("roundrobin", &RoundRobinBalance{})
}

// 轮询算法
func (p *RoundRobinBalance) DoBalance(insts []*Instance, key ...string) (*Instance, error) {
	lens := len(insts)
	if lens == 0 {
		return nil, fmt.Errorf("No Instance found")
	}

	if p.curIndex >= lens {
		p.curIndex = 0
	}
	// 获取实例
	inst := insts[p.curIndex]
	// 计算下一个实例的索引值
	p.curIndex = (p.curIndex + 1) % lens

	inst.CallTimes++

	return inst, nil
}
