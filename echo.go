package main

import (
	"fmt"
	"github.com/grokify/html-strip-tags-go"
	"github.com/microcosm-cc/bluemonday"
	"hash/crc32"
	"sort"
)

func main() {
	p := bluemonday.UGCPolicy()

	// The policy can then be used to sanitize lots of input and it is safe to use the policy in multiple goroutines
	html := p.Sanitize(
		//`<a onblur="alert(secret)" href="http://www.google.com">Google</a>`,
		`Period 3 Letter time &amp; Rhyme time`,
	)
	stripped := strip.StripTags(html)

	// Output:
	// <a href="http://www.google.com" rel="nofollow">Google</a>
	fmt.Println(html)
	fmt.Println(stripped)

	sli := []int{0, 1, 2, 3}
	var reMap map[int]Re
	fmt.Println("---------------: ", reMap[1])
	t := sli[:0]
	fmt.Println("---------------sli: ", t)

	//检索
	a := []float64{1.0, 2.0, 3.3, 4.6, 6.1, 7.2, 8.0}
	x := 1.0
	i := sort.SearchFloat64s(a, x)
	fmt.Printf("found %g at index %d in %v\n", x, i, a) //found 2 at index 1 in [1 2 3.3 4.6 6.1 7.2 8] 如果不存在返回0    a := []int{1, 2, 3, 4, 6, 7, 8}
	b := []int{1, 2, 3, 4, 6, 7, 8}
	y := 2
	j := sort.SearchInts(b, y)
	fmt.Printf("found %d at index %d in %v\n", y, j, b) //found 2 at index 1 in [1 2 3 4 6 7 8]

	z := 1
	i = sort.SearchInts(b, z)
	fmt.Printf("%d not found, can be inserted at index %d in %v\n", z, i, b) //1 not found, can be inserted at index 0 in [1 2 3 4 6 7 8]

}

// 计算两个数的最大公约数
func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func hashValue(origin string) {
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(origin), crcTable)
	fmt.Println("hashValue-----------", hashVal)
}

type Re struct {
	CourseId int `json:"course_id"`
	*Base
}

type Base struct {
	Grade   int `json:"grade"`
	Subject int `json:"subject"`
}

// 64位平台，对齐参数是8
type User struct {
	A int32 // 4
	//	B []int32 // 24
	C string // 16
	D bool   // 1
}
