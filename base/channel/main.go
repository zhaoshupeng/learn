package main

import (
	"fmt"
)

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func main1() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			data, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- data * data
		}
		close(ch2)
	}()
	for v := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println("res: ", v)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go counter(ch1)
	go square(ch2, ch1)
	printer(ch2)
}

func counter(out chan<- int) {
	for i := 0; i <= 9; i++ {
		out <- i
	}
	close(out)
}

func square(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println("---------: ", v)
	}
}
