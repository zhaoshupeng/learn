package balance

import (
	"fmt"
)

// https://www.cnblogs.com/wsw-seu/p/11336634.html
// https://my.oschina.net/doocs/blog/4486813

//WeightRoundRobin负载均衡
type WeightRoundRobinBalance struct {
	Index  int64 // 表示本次请求到来时，选择的服务器索引，初始值为 0 (有的文章中普遍初始值是 -1)
	Weight int64 //表示当前调度的权值，初始值为max(S),即服务器最大的权重数
}

func init() {
	RegisterBalancer("weight_roundrobin", &WeightRoundRobinBalance{})
}

func (p *WeightRoundRobinBalance) DoBalance(insts []*Instance, key ...string) (*Instance, error) {
	lens := len(insts)
	if lens == 0 {
		return nil, fmt.Errorf("No Instance found")
	}

	inst := p.GetInst(insts)
	inst.CallTimes++
	fmt.Println("普通加权轮询", inst.Host, inst.Weight, inst.CallTimes)

	return inst, nil
}

// 优化方法：平滑的加权轮询

/**
轮询加权重负载策略(普通):
	设计一个权重因子，初始值为所有被调用的结点中最大权重值。 负载均衡使用轮询算法，被选中结点权重值大于等于权重因子则可以调用，
否则用下一结点的权重值与权重因子比较，一轮循环结束后如果没有选中结点，则降低权重因子，继续通过与权重因子比较进行选择，直到选中为止。
权重因子降为 0 后，恢复为最大权重值。

*/
/**
这背后的数学原理，自己思考了一下，总结如下：
　　current_weight的值，其变化序列就是一个等差序列：max, max-gcd, max-2gcd, …, 0(max)，
将current_weight从max变为0的过程，称为一个轮回。
　　针对每个current_weight，该算法就是要把服务器数组从头到尾扫描一遍，将其中权重大于等于current_weight的所
有服务器填充到结果序列中。扫描完一遍服务器数组之后，将current_weight变为下一个值，再一次从头到尾扫描服务器数组。
　　在current_weight变化过程中，不管current_weight当前为何值，具有max权重的服务器每次肯定会被选中。因此，
具有max权重的服务器会在序列中出现max/gcd次（等差序列中的项数）。
　　更一般的，当current_weight变为x之后，权重为x的服务器，在current_weight接下来的变化过程中，每次都会被选中，因此，具有x权重的服务器，会在序列中出现x/gcd次。所以，每个服务器在结果序列中出现的次数，是与其权重成正比的，这就是符合加权轮询算法的要求了。
*/
func (p *WeightRoundRobinBalance) GetInst(insts []*Instance) *Instance {
	//计算多个数的最大公约数(步长)
	gcd := getGCD(insts)
	for {
		p.Index = (p.Index + 1) % int64(len(insts))
		// 已全部被遍历完一次
		if p.Index == 0 {
			p.Weight = p.Weight - gcd // 权重因子每次降低的步长;权重因子的降低步长为所有结点权重值的最大公约数。
			// 赋值currentWeight为0,回归到初始状态
			if p.Weight <= 0 {
				p.Weight = getMaxWeight(insts) //获取最大权重
				if p.Weight == 0 {
					return &Instance{}
				}
			}
		}

		// 直到当前遍历实例的weight大于或等于currentWeight
		if insts[p.Index].Weight >= p.Weight {
			return insts[p.Index]
		}
	}
}

// todo understand
//计算两个数的最大公约数
func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

//计算多个数的最大公约数
func getGCD(insts []*Instance) int64 {
	var weights []int64

	for _, inst := range insts {
		weights = append(weights, inst.Weight)
	}

	g := weights[0]
	for i := 1; i < len(weights)-1; i++ {
		oldGcd := g
		g = gcd(oldGcd, weights[i])
	}
	return g
}

//获取最大权重
func getMaxWeight(insts []*Instance) int64 {
	var max int64 = 0
	for _, inst := range insts {
		if inst.Weight >= int64(max) {
			max = inst.Weight
		}
	}

	return max
}
