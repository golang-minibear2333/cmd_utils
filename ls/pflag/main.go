package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
)

func ShowPath(targetDirPath string, isAll bool) {
	//我其实也没找到怎么获取隐藏文件的办法，只能出此下策暂时弄个假的
	if isAll{
		fmt.Print(". .. .git ")
	}
	if dirList, err := os.ReadDir(targetDirPath); err == nil {
		for _, dirInfo := range dirList {
			fmt.Print(dirInfo.Name() + " ")
		}
		fmt.Println()
	} else {
		fmt.Println(err.Error())
	}
}
func main() {
	var a = pflag.BoolP("all", "a", false, "Include directory entries whose names begin with a dot (.).")
	var help = pflag.BoolP("help", "h", false, "Show this help message.")
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}
	args := pflag.Args()
	if len(args) == 0 {
		args = append(args, "./")
	}
	if len(args) == 1 {
		ShowPath(args[0], *a)
	} else {
		for _, v := range args {
			fmt.Println(v + ":")
			ShowPath(v, *a)
		}
	}
}
