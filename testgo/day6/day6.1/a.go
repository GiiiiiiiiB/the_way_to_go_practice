package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"unicode/utf8"
)

func t1() {
	var a = [5]int{1, 2, 3, 4, 5}
	c := a
	b := a[:]
	d := a[0:]
	e := a[1 : len(a)-2]
	b[2] = 5

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func t2() {
	a := make([]int, 5, 10)
	for _, j := range a {
		fmt.Println(j)
	}
}

func t3() {
	s := make([]byte, 5)
	fmt.Println(len(s), cap(s))
	s = s[2:4]
	fmt.Println(len(s), cap(s))
}

func Sum(arrF []float32) float64 {
	s := 0.0
	for _, v := range arrF {
		s += float64(v)
	}
	return s
}

func t4() {
	arr := [...]int{3, 1, 15, 3, 1, 3, 4, 5, 6, 7, 9}
	arrF := arr[:]
	b := arrF[3:3]
	fmt.Println(b, len(b), cap(b))
}

func t5() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 2)

	n := copy(slTo, slFrom)
	fmt.Println(&n, &slTo[0])
	slFrom[0] = 2
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	fmt.Println(len(sl3), cap(sl3))

	sl4 := append(sl3, 4, 5, 6)
	sl4[0] = 2
	fmt.Println(sl3)
}

func t6(data ...int) {
	fmt.Printf("%T\n", data)
	for _, v := range data {
		fmt.Println(v)
	}
	fmt.Println("-------------")
}

func t7(data []int) {
	fmt.Printf("%T\n", data)
	for _, v := range data {
		fmt.Println(v)
	}
	fmt.Println("-------------")
}

func t8(s []int, factor int) []int {
	return make([]int, len(s)*factor)
}

func t9(a int) bool {
	if n := a % 2; n == 0 {
		return true
	}
	return false
}

func t10(s [10]int, f func(a int) bool) []int {
	l0 := make([]int, 10)
	ix := 0
	for _, v := range s {
		if f(v) {
			l0[ix] = v
			ix++
		}
	}
	return l0
}

func test_t10() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(t10(a, t9))
}

func t_11() {
	a := []int{1, 2, 3, 4, 45, 65, 5, 7}

	fmt.Println(sort.SearchInts(a, 5))
	fmt.Println(a)

	fmt.Println(sort.IntsAreSorted(a))
	sort.Ints(a)
	fmt.Println(a)
	fmt.Println(sort.IntsAreSorted(a))
}

func t_12() {
	var digitRegexp = regexp.MustCompile("[0-9]+")

	fileBytes, _ := ioutil.ReadFile("./a.txt")
	fmt.Println(fileBytes)
	fmt.Println("------------")
	b := digitRegexp.FindAll(fileBytes, len(fileBytes))
	fmt.Println(b)
	c := make([]byte, 0)
	for _, bytes := range b {
		c = append(c, bytes...)
	}
	for _, v := range c {
		fmt.Printf("%c\n", v)
	}
}

func t_13(str string, i int) (string, string) {
	return str[:i], str[i:]
}

func t_13_2(str string) string {
	return str[len(str)/2:] + str[:len(str)/2]
}

func t_14() {
	a := "Google"
	e := a[:]
	fmt.Printf("%T\n", e)
	fmt.Println(e)
	c := make([]byte, len(a))
	copy(c, a[:])
	d := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		d[len(a)-1-i] = c[i]
	}
	for i := 0; i < len(c)/2; i++ {
		tmp := c[len(c)-1-i]
		c[len(c)-1-i] = c[i]
		c[i] = tmp
	}
	fmt.Println(string(d))
	fmt.Println(string(c))
	// fmt.Println(strings.Join(d, ""))
}

func t_15() {
	a := "google支持中文吗"
	b := make([]rune, utf8.RuneCountInString(a))
	b = []rune(a)
	for i := 0; i < len(b)/2; i++ {
		tmp := b[len(b)-i-1]
		b[len(b)-i-1] = b[i]
		b[i] = tmp
	}
	fmt.Println(string(b))
}

func t_16(arr []int) {
	a := make([]int, len(arr))
	ix := 0
	tmp := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] != tmp {
			a[ix] = arr[i]
			ix++
		}
		tmp = arr[i]
	}
	fmt.Println(a)

}

func test_16() {
	a := []int{1, 2, 2, 3, 4, 4, 5}
	t_16(a)
}

func t_17() {
	a := []int{1, 2, 7, 8, 3, 4, 5}
	for i := 0; i < len(a)-1; i++ {
		for g := 0; g < len(a)-1-i; g++ {
			if j, k := a[g], a[g+1]; j > k {
				a[g+1] = j
				a[g] = k
			}
		}

	}
	fmt.Println(a)
}

func xx(a int) int {
	return a * 10
}

func t_18(f func(a int) int, arr []int) []int {
	a := make([]int, len(arr))
	for ix, v := range arr {
		a[ix] = f(v)
	}
	fmt.Println(a)
	return a
}

func main() {
	a := []int{1, 3, 5, 7, 9}
	t_18(xx, a)
}
