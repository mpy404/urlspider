package main

import (
	"URLScan/scan"
	"fmt"
)

func run(method, googleHack string) {
	if method == "baidu" {
		scan.BaiduRun(googleHack)
	} else if method == "bing" {
		scan.BingRun(googleHack)
	}
}

func input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return text
}

func main() {
	method := input("请输入搜索引擎[baidu、bing] : ")
	googleHack := input("请输入搜索语法[inurl:php?id=1+公司](注意:空格使用+代替) : ")
	run(method, googleHack)
}
