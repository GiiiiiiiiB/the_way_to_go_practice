package main

import (
	"fmt"
)

func t2(a []int) {
	a[1] = 5
	for i, j := range a {
		fmt.Println(i, j)
	}
}

func t1() {
	// var arr1 [5]int

	// for i := 0; i < len(arr1); i++ {
	// 	arr1[i] = i * 2
	// }
	var arr1 = []int{1, 2, 3, 4, 45}

	t2(arr1)

	for i, j := range arr1 {
		fmt.Println(i, j)
	}

	for i := range arr1 {
		fmt.Println(i)
	}
}

func fp(a *[3]int) { fmt.Println(a[1]) }

func t3() {
	x := []int{2, 3, 4, 5, 6}
	x[4] = 10
	fmt.Println(len(x), cap(x))
	y := x[1:3]
	// y[2] = 4
	fmt.Println(len(y), cap(y))
}

func main() {
	t1()
	t3()
}
