package balance

// 负载均衡器
type Balancer interface {
	DoBalance([]*Instance, ...string) (*Instance, error)
}
