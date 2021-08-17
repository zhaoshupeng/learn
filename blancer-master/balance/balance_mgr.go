package balance

import "fmt"

var (
	mgr = BalanceMgr{
		allBalance: make(map[string]Balancer),
	}
)

type BalanceMgr struct {
	allBalance map[string]Balancer // key 是对应算法的名称, value 对应的是负载均衡器
}

// 注册特定的负载均衡算法
func (p *BalanceMgr) registerBalancer(balanceType string, b Balancer) {
	p.allBalance[balanceType] = b
}

func RegisterBalancer(balanceType string, b Balancer) {
	mgr.registerBalancer(balanceType, b)
}

// 执行特定的负载均衡算法
func DoBalance(balanceType string, insts []*Instance) (*Instance, error) {
	Balancer, ok := mgr.allBalance[balanceType]
	if !ok {
		fmt.Printf("Not found %s balancer", balanceType)
		return nil, fmt.Errorf("not fund balancer")
	}
	return Balancer.DoBalance(insts)
}
