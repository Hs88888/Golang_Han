package main

import (
	"fmt"
	"os"
	"strconv"

	//"/src/GoProgramming/userA/tempconv"			// 包导入问题
)

func main() {
	fmt.Println("----test----")
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(t)
		//f := tempconv.Fahrenheit(t)
		//c := tempconv.Celsius(t)
		//fmt.Printf(%s = %s, %s = %s\n, f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
