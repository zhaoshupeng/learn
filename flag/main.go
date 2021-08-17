package main

import (
	"flag"
	"fmt"
)

var flagName1 = flag.Int("flagname", 1234, "help message for flagname")

var flagvar int

func init() {
	flag.IntVar(&flagvar, "var", 55, "help message for var")
}

func main() {
	// 解析命令行参数到定义的flag
	// 在main方法中调用flag.Parse从os.Args[1:]中解析选项。因为os.Args[0]为可执行程序路径，会被剔除。
	// flag.Parse方法必须在所有选项都定义之后调用，且flag.Parse调用之后不能再定义选项。
	flag.Parse()
	// 遇到第一个非选项参数（即不是以-和--开头的）或终止符--，解析停止。
	// go run main.go noflag -flagname 5556 //因为解析遇到noflag就停止了，后面的选项-intflag没有被解析到。所以所有选项都取的默认值。
	// 解析终止之后如果还有命令行参数，flag库会存储下来，通过flag.Args方法返回这些参数的切片。可以通过flag.NArg方法获取未解析的参数数量，flag.Arg(i)访问位置i（从 0 开始）

	fmt.Println("ip has value ", *flagName1)
	fmt.Println("flagvar has value ", flagvar)

	for i := 0; i < flag.NArg(); i++ {
		fmt.Printf("Argument %d: %s\n", i, flag.Arg(i))
	}

	fmt.Println(flag.Args())
	fmt.Println(flag.Arg(0))

}
