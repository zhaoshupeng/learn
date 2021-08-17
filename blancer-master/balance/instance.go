package balance

import "strconv"

// 一个实例的信息:模拟后端真是服务器
type Instance struct {
	Host      string
	Port      int64
	Weight    int64
	CallTimes int64
	Type      string
}

// 创建一个实例
func NewInstance(host string, port int64, wc int64) *Instance {
	return &Instance{
		Host:      host,
		Port:      port,
		Weight:    wc, // 权重
		CallTimes: 0,
	}
}

func (i *Instance) GetPort() int64 {
	return i.Port
}

func (i *Instance) GetHost() string {
	return i.Host
}

func (i *Instance) GetCallTimes() int64 {
	return i.CallTimes
}

func (i *Instance) GetResult() string {
	return i.Host + ":" + strconv.FormatInt(i.Port, 10) + ";call times: " + strconv.FormatInt(i.CallTimes, 10)
}
