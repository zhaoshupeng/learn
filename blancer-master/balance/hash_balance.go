package balance

import (
	"fmt"
	"hash/crc32"
	"math/rand"
)

// https://segmentfault.com/a/1190000021199728  //一致性哈希算法
// https://learnku.com/articles/30269   PHP 之一致性 hash 算法

// https://zhuanlan.zhihu.com/p/61636624  CRC（循环冗余校验码）简介与实现解析

// 一致性hash负载均衡
type HashBalance struct {
}

func init() {
	RegisterBalancer("hash", &HashBalance{})
}

// 不带虚拟节点的一致性哈希
func (p *HashBalance) DoBalance(insts []*Instance, key ...string) (*Instance, error) {
	var defKey string = fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		defKey = key[0]
	}

	lens := len(insts)
	if lens == 0 {
		return nil, fmt.Errorf("No backend instance")
	}

	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % lens
	inst := insts[index]
	inst.CallTimes++

	return inst, nil
}
