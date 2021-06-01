package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func testTime() {
	a := time.Now()
	fmt.Println(a)
	fmt.Println(a.Day())
	fmt.Println(a.UnixNano())
	fmt.Println(a.Format("2006-05-02 10:31:38"))
	fmt.Println(a.Format(time.ANSIC))
	fmt.Println(a.Format(time.RFC822))
	fmt.Println(a.UTC())
	fmt.Println(a.Unix())

	week := 7 * 24 * 60 * 60 * 1e9
	week_from_new := a.Add(time.Duration(week))
	fmt.Println(week_from_new)
}

func testP() {
	var l1 = 5
	var intP *int
	intP = &l1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
	l1 = 6
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
	*intP++
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}

func test11() {
	math.Abs(-3)
}

func test() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

func Season(i int) string {

	switch i {
	case 1, 2, 3:
		return "chunji"
	case 4, 5, 6:
		return "xiatian"
	case 7, 8, 9:
		fmt.Println("专利之星吗")
		break
		return "qiutian"
	case 10, 11, 12:
		return "dongtian"
	default:
		return "error"
	}
	fmt.Println("hahahhahahhahahah")
	return "error"
}

func testLoop() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}
	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}
}

func testLoop2() {
	for i := 0; i < 15; i++ {
		fmt.Println(i)
	}

	j := 0
lp:
	if j > 15 {
		return
	} else {
		fmt.Println(j)
		j++
		goto lp
	}

}

func test2() {
	for i := 1; i <= 6; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}

	for i := 1; i <= 6; i++ {
		fmt.Println(strings.Repeat("G", i))
	}
}

func test3() {
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
}

func test4() {
	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}

func main() {
	fmt.Print(Season(7))

}
