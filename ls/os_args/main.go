package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//使用 os.Agrs 来获取传入程序的参数
//使用 ioutil.ReadDir 读取目录,将返回值[] FileInfo遍历,输出目录/文件信息
//将遇到的error 输出到 终端
func main() {
	targetDirPath := "./"
	if len(os.Args) > 1 {
		targetDirPath = os.Args[1]
	}
	if dirList, err := ioutil.ReadDir(targetDirPath); err == nil {
		for _, dirInfo := range dirList {
			fmt.Print(dirInfo.Name() + " ")
		}
	} else {
		fmt.Println(err.Error())
	}
}
