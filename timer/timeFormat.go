package main

import (
	"fmt"
	"strconv"
	"time"
)

var MinuteDataMap = make(map[int]map[string]map[string]int, 0) //[time:topic:event]:num

func main() {
	tm := time.Unix(1622446700, 0).Format("2006-01-02 15:04:05")
	fmt.Println("wwwwwwwwwwwwwwww", tm)

	// 格式化天时间戳
	timeObj, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)

	todayBeginUnix := timeObj.Unix()
	twoDaysAfterBeginUnix := todayBeginUnix + 2*86400
	sevenDaysAgoUnix := todayBeginUnix - 7*86400

	//str := "ALTER TABLE accounting ADD PARTITION (PARTITION p" + string(time.Unix(todayBeginUnix, 0).Format("20060102")) + " VALUES LESS THAN (" + string(todayBeginUnix) + "));"

	//todayInt := time.Unix(todayBeginUnix, 0).Format("20060102")
	todayStr := string(time.Unix(todayBeginUnix, 0).Format("20060102"))
	fmt.Println(todayStr)

	twoDaysAfterSQL := "ALTER TABLE accounting ADD PARTITION (PARTITION p" + string(time.Unix(twoDaysAfterBeginUnix, 0).Format("20060102")) + " VALUES LESS THAN (" + strconv.FormatInt(twoDaysAfterBeginUnix, 10) + "));"

	sevenDaysAgoSQL := "ALTER TABLE accounting DROP PARTITION p" + string(time.Unix(sevenDaysAgoUnix, 0).Format("20060102")) + ";"

	fmt.Println("twoDaysAfterSQL: ", twoDaysAfterSQL)
	fmt.Println("sevenDaysAgoSQL: ", sevenDaysAgoSQL)

	fmt.Println("todayBeginUnix----", time.Unix(todayBeginUnix, 0).Format("20060102"))
	fmt.Println("twoDaysAfterBeginUnix----", time.Unix(twoDaysAfterBeginUnix, 0).Format("20060102"))
	fmt.Println("sevenDaysAgoUnix----", time.Unix(sevenDaysAgoUnix, 0).Format("20060102"))

	//fmt.Println(timeObj.Unix() + 86400)
	//fmt.Println(timeObj.Format("20060102"))
	//time1 := time.Now().Unix()
	//fmt.Println("time1-----", time1)

	//minuteObj1, _ := time.ParseInLocation("2006-01-02 15:04:00", time.Now().Format("2006-01-02 15:04:00"), time.Local)
	//minute1 := minuteObj1.Unix()
	// 格式化分时间戳
	minuteObj1, _ := time.ParseInLocation("2006-01-02 15:04:00", time.Unix(1616487231, 0).Format("2006-01-02 15:04:00"), time.Local)
	minute1 := minuteObj1.Unix()

	fmt.Println("minute: ", minute1)

	//CacheTopicMinuteData("firtopic", "publish")
	//CacheTopicMinuteData("firtopic", "publish")
	//CacheTopicMinuteData("firtopic", "consume")
	//CacheTopicMinuteData("secondtopic", "consume")
	//fmt.Println(MinuteDataMap)
}
func CacheTopicMinuteData(topic, event string) {
	minuteObj, _ := time.ParseInLocation("2006-01-02 15:04:00", time.Now().Format("2006-01-02 15:04:00"), time.Local)
	minute := minuteObj.Unix()
	if _, ok := MinuteDataMap[int(minute)][topic][event]; ok {
		MinuteDataMap[int(minute)][topic][event] += 1 // 存在值
	} else {
		fmt.Println(11111)
		//MinuteDataMap[int(minute)] = map[string]map[string]int{topic: {event: 1}}

		//MinuteDataMap[int(minute)][topic][event] = 1

		if _, ok := MinuteDataMap[int(minute)][topic]; ok {
			//MinuteDataMap[int(minute)][topic][event] = 1
			MinuteDataMap[int(minute)][topic][event] = 1

			return
		} else {
			if _, ok := MinuteDataMap[int(minute)]; ok {
				//MinuteDataMap[int(minute)][topic][event] = 1
				MinuteDataMap[int(minute)][topic] = map[string]int{event: 1}

				return
			} else {
				MinuteDataMap[int(minute)] = map[string]map[string]int{topic: {event: 1}}
			}
		}

	}
}
