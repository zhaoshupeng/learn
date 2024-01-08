package main

import (
	"fmt"
	"runtime"
	"time"
)

func main1() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
		time.Sleep(10 * time.Microsecond)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
		time.Sleep(10 * time.Microsecond)
	}
}

func main() {
	fmt.Println("---------------: ", (3-4)/2)

	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}
