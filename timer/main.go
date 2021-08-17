package main

import (
	"fmt"
	"time"
)

func main1() {
	t := time.NewTimer(time.Second * 2)
	// 同时需要注意 defer t.Stop()在这里并不会停止定时器。这是因为Stop会停止Timer，停止后，Timer不会再被发送，但是Stop不会关闭通道，防止读取通道发生错误。
	//如果想停止定时器，只能让go程序自动结束。
	defer t.Stop()
	for {
		<-t.C
		timeStr := time.Now().Format("2006-01-02 15:04:05")
		fmt.Println("timer running...", timeStr)
		// 需要重置Reset 使 t 重新开始计时,间隔时间5s
		// 注释 t.Reset(time.Second * 5)会导致通道堵塞，报fatal error: all goroutines are asleep - deadlock!错误。
		t.Reset(time.Second * 5)
	}
}

//如果想停止定时器，只能让go程序自动结束。
func main() {
	t := time.NewTimer(time.Second * 3)
	t1 := time.NewTimer(time.Second * 3)

	ch := make(chan bool)
	go func(t *time.Timer) {
		defer t.Stop()
		defer t1.Stop()

		for {
			select {
			case <-t.C:
				timeStr := time.Now().Format("2006-01-02 15:04:05")
				fmt.Println("timer running....", timeStr)
				// 需要重置Reset 使 t 重新开始计时
				t.Reset(time.Second * 3)
			case <-t1.C:
				timeStr := time.Now().Format("2006-01-02 15:04:05")
				fmt.Println("timer running1....", timeStr)
				// 需要重置Reset 使 t 重新开始计时
				t.Reset(time.Second * 3)
			case stop := <-ch:
				if stop {
					fmt.Println("timer Stop")
					return
				}
			default:
				fmt.Println("default....")
			}
			fmt.Println("111111")
		}
	}(t)
	time.Sleep(20 * time.Second)
	ch <- true
	close(ch)
	time.Sleep(1 * time.Second)
}

func main21() {
	t := time.NewTicker(time.Second * 2)
	// 这里的defer t.Stop()和上面示例相似，也不会停止定时器，解决办法一样。
	defer t.Stop()
	for {
		<-t.C
		timeUnix := time.Now().Unix()
		fmt.Println("Ticker running...", timeUnix)
	}
}

func main12() {

	ticker := time.NewTicker(2 * time.Second)
	ticker1 := time.NewTicker(2 * time.Second)

	ch := make(chan bool)
	go func(ticker, ticker1 *time.Ticker) {
		defer ticker.Stop()
		defer ticker1.Stop()
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-ticker.C:
				time.Sleep(4 * time.Second)

				fmt.Println("ticker", time.Now().Format("2006-01-02 15:04-05"))
				//fmt.Println("Ticker running...")
			case <-ticker1.C:
				time.Sleep(4 * time.Second)

				fmt.Println("ticker1", time.Now().Format("2006-01-02 15:04-05"))
				//fmt.Println("Ticker1 running...")
			case stop := <-ch:
				if stop {
					fmt.Println("Ticker Stop")
					return
				}
			}
		}
	}(ticker, ticker1)
	time.Sleep(10 * time.Second)
	ch <- true
	close(ch)
}

func main3() {
	t := time.After(time.Second * 3)
	fmt.Printf("t type=%T\n", t)
	//阻塞3秒
	fmt.Println("t=", <-t)
}

// select语句阻塞等待最先返回数据的channel`,如ch1通道成功读取数据，则先输出1th case is selected. e1=1，之后每隔2s输出 Timed out。
func main31() {
	ch1 := make(chan int, 1)
	ch1 <- 1
	for {
		select {
		case e1 := <-ch1:
			//如果ch1通道成功读取数据，则执行该case处理语句
			fmt.Printf("1th case is selected. e1=%v\n", e1)
		// 通过源码我们发现它返回的是一个NewTimer(d).C，其底层是用NewTimer实现的，所以如果考虑到效率低，可以直接自己调用NewTimer。
		// time.After()表示多长时间长的时候后返回一条time.Time类型的通道消息。但是在取出channel内容之前不阻塞，后续程序可以继续执行。
		case t := <-time.After(time.Second * 2):
			fmt.Println("Timed out", t)
		}
	}

}
