package main

import (
	"fmt"
	"time"
)

var m = make(map[int]int)

func main() {
	//一个go程写map
	go func() {
		for i := 0; i < 10000; i++ {
			m[i] = i
		}
	}()

	time.Sleep(time.Second * 5)

	//一个go程读map
	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Println(m[i])
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Println(m[i])
		}
	}()
	time.Sleep(time.Second * 20)
}
