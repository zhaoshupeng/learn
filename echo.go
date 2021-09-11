package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash/crc32"
)

func main() {

	v1 := 3 / 4
	v2 := 3 / 4 * 4
	fmt.Println("----------------------", v1, v2)
	//md5Str("good")
	//hashValue("good")
	//buckets := make([][]int, 3)
	//fmt.Println(len(buckets), buckets)
	//
	//arr := []int{5, 6, 8, 9}
	//for val := range arr {
	//	fmt.Println("val----", val)
	//}
	//
	//var num int64 = 10
	//fmt.Printf("变量类型: %T; 变量占用的字节数: %d; 变量地址：%p\n", num, unsafe.Sizeof(num), &num)
	//
	//var sli []interface{}
	//fmt.Printf("变量类型: %T; 变量占用的字节数: %d; 变量地址：%p; 变量值:%+v\n", sli, unsafe.Sizeof(sli), &sli, sli)
	//fmt.Println(sli)
	//fmt.Println(gcd(8, 3))
}

//计算两个数的最大公约数
func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func md5Str(origin string) {

	m := md5.New()
	m.Write([]byte(origin))

	str := hex.EncodeToString(m.Sum(nil))
	fmt.Println("md5Str---------", str)
	//return hex.EncodeToString(m.Sum(nil))

	TestString := "Hi, pandaman!"

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestString))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)
}

func hashValue(origin string) {
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(origin), crcTable)
	fmt.Println("hashValue-----------", hashVal)
}
