/*
	2.3 指针
 */
//package main
//
//import "fmt"
//
//func main() {
//	v := 1
//	fmt.Println(v)
//	fmt.Println(f(&v))
//}
//
//func f(p *int) int {
//	*p++
//	return *p
//}

/*
	2.3 指针（命令行参数）
 */
//package main
//
//import (
//	"flag"
//	"fmt"
//	"strings"
//)
//
//var n = flag.Bool("n", false, "omit trailing newline")
//var sep = flag.String("s", " ", "separator")
//
//func main() {
//	flag.Parse()
//	fmt.Print(strings.Join(flag.Args(), *sep))
//	if !*n {
//		fmt.Println()
//	}
//}

package main

import "fmt"

func main() {
	p := new(int)			// 表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
}
