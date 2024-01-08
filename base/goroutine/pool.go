package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	job *Job
	sum int
}

func createPool(num int, jobChan chan *Job, result chan *Result) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, result chan *Result) {
			for job := range jobChan {
				r_num := job.RandNum

				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num = r_num / 10
				}
				res := &Result{
					job: job,
					sum: sum,
				}
				result <- res
			}
		}(jobChan, result)
	}
}

func main() {
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)
	// 创建协程池
	createPool(64, jobChan, resultChan)
	// 打印结果
	go func(resultChan chan *Result) {
		for res := range resultChan {
			fmt.Println("----------job id: ", res.job.Id, "job value:", res.job.RandNum, "sum: ", res.sum)
		}
	}(resultChan)

	// 产生任务
	var id int
	for {
		job := &Job{
			Id:      id,
			RandNum: rand.Int(),
		}
		jobChan <- job
		id++
		time.Sleep(500 * time.Millisecond)
	}
}
