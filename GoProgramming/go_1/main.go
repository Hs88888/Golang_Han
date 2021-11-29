/*
	1.2 命令行参数
 */
//package main
//
//import (
//	"fmt"
//	"os"
//)
//
//func main() {
//	var s, sep string
//	for i := 1; i < len(os.Args); i++ { 	//os.Args的第一个参数是命令本身的名字，其它的元素则是程序启动时传给它的参数
//		s += sep + os.Args[i]
//		sep = " "
//	}
//	fmt.Println(s)
//}

/*
	1.3 查找重复的行
 */
//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//func main() {
//	counts := make(map[string]int)
//	input := bufio.NewScanner(os.Stdin)
//	for i := 1; i <= 5; i++ {
//		input.Scan()
//		counts[input.Text()]++
//	}
//
//	for line, n := range counts {
//		if n > 1 {
//			fmt.Printf("%d\t%s\n", n, line)
//		}
//	}
//}

/*
	1.3 查找重复的行（从文件中读取数据）
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d%t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
