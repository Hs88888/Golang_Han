package main

import "fmt"

type Saiyan struct {
	Name string
	Power int
}

func (s *Saiyan) Super() {
	s.Power += 1000
}

//func Super(s *Saiyan) {
//	//s.Power += 10000
//	s = &Saiyan{"Gohan", 1000}
//}

func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1
	source[index], source[lastIndex] = source[lastIndex], source[index]
	return source[:lastIndex]
}

func main() {
	//power := 9000
	//fmt.Printf("It's over %d\n", power)

	//if len(os.Args) != 2 {
	//	os.Exit(1)
	//}
	//fmt.Println("It's over", os.Args[1])

	//goku := &Saiyan{"Power", 9000}
	//Super(goku)
	//fmt.Println(goku.Power)

	//goku := &Saiyan{"Goku", 9000}
	//goku.Super()
	//fmt.Println(goku.Power)

	//scores := make([]int, 5)  // 长度、容量
	//scores[0] = 1
	//scores = append(scores, 10)
	//fmt.Println(scores)
	//c := cap(scores)
	//fmt.Println(c)		//5
	//
	//for i := 0; i < 25; i++ {
	//	scores = append(scores, i)
	//
	//	if cap(scores) != c {
	//		c = cap(scores)
	//		fmt.Println(c)
	//	}
	//}

	scores := []int{1, 2, 3, 4, 5}
	//slice := scores[2:4]
	//slice[0] = 999

	scores = removeAtIndex(scores, 2)
	fmt.Println(scores)
}