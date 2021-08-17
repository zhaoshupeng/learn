package main

import (
	"fmt"
	"sync/atomic"
)

/**
1.import "sync/atomic"
	atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用。
	这些函数必须谨慎地保证正确使用。除了某些特殊的底层应用，使用通道或者sync包的函数/类型实现同步更好。

2.goalng 中的原子操作类型
　　int32、int64、uint32、uint64、uintptr和unsafe.Pointer类型，共6个

3.CAS操作的优势是，可以在不形成临界区和创建互斥量的情况下完成并发安全的值替换操作。
	这可以大大的减少同步对程序性能的损耗。
	当然，CAS操作也有劣势。在被操作值被频繁变更的情况下，CAS操作并不那么容易成功。

https://www.jianshu.com/p/a0be632df99b
https://studygolang.com/articles/3557
*/

var value int32

func main() {
	// 访问量计数
	var count int64 = 0
	// 对count变量进行原子加 1
	// 原子操作可以在并发环境安全的执行
	atomic.AddInt64(&count, 1)

	// 对count变量原子减去10
	atomic.AddInt64(&count, -10)

	// 原子读取count变量的内容
	pv := atomic.LoadInt64(&count)
	fmt.Println(pv)

	var Atomicvalue atomic.Value
	Atomicvalue.Store([]int{1, 2, 3, 4, 5})
	anotherStore(Atomicvalue)
	fmt.Println("main: ", Atomicvalue)

}
func anotherStore(Atomicvalue atomic.Value) {
	Atomicvalue.Store([]int{6, 7, 8, 9, 10})
	fmt.Println("anotherStore: ", Atomicvalue)
}

// CAS的使用示例
// 由示例可以看出，我们需要多次使用for循环来判断该值是否已被更改，为了保证CAS操作成功，仅在 CompareAndSwapInt32 返回为 true时才退出循环，这跟自旋锁的自旋行为相似。
func AddValue(delta int32) {
	for {
		v := value
		if atomic.CompareAndSwapInt32(&value, v, (v + delta)) {
			break
		}
	}
}
