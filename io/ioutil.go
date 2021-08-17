package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

/**
1.在go语言中文件读写相关的包有：
	io/ioutil - 提供常用的io函数，例如：便捷的文件读写。
	os - 系统相关的函数，我这里主要用到跟文件相关的函数，例如，打开文件。
	path/filepath - 用处理文件路径

https://www.tizi365.com/archives/337.html
*/

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// 1. 读文件
	// ReadFile 从filename指定的文件中读取数据并返回文件的内容。成功的调用返回的err为nil而非EOF。因为本函数定义为读取整个文件，它不会将读取返回的EOF视为应报告的错误。
	content, err := ioutil.ReadFile("./print.go")
	if err != nil {
		// 读取文件失败
		panic(err)
	}
	fmt.Println(string(content))

	// 2. 写文件
	contents := "这里是文件的内容。"
	// 因为WriteFile函数接受的文件内容是字节数组，所以需要将content转换成字节数组
	// 0666是指文件的权限是具有读写权限，具体可以参考linux文件权限相关内容。
	err1 := ioutil.WriteFile("./demo.txt", []byte(contents), 0666)

	if err1 != nil {
		panic(err)
	}

	// 3. 检测文件是否存在
	_, err = os.Stat("./demo.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
			return
		}

		if os.IsPermission(err) {
			fmt.Println("没有权限对文件进行操作。")
			return
		}

		fmt.Println("其他错误。")
		return
	}

	// err == nil 则表示文件存在
	fmt.Println("文件存在")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	// 4.获取文件名字
	filename := filepath.Base("/images/logo.jpg")
	fmt.Println(filename) // logo.jpg

	// 获取目录
	dir := filepath.Dir("/images/logo.jpg")
	fmt.Println(dir) // \images
	dir1 := filepath.Dir("./demo.txt")
	fmt.Println(dir1) // .

}
