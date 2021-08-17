package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	c := make(map[string]string)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			k, v := strconv.Itoa(n), strconv.Itoa(n)
			c[k] = v
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(c)
}
