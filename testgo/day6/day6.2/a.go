package main

import (
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
	// "sync"
	// "syscall"
)

func t1() {
	mapLit := map[string]int{"one": 1, "two": 2}
	fmt.Println(&mapLit)
}

func t2() {
	// Version A:
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	// Version B: NOT GOOD!
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1) // item is only a copy of the slice element.
		item[1] = 2                 // This 'item' will be lost on the next iteration.
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)
}

func t3() {
	a := []int{1, 34, 5, 6, 9}
	for _, v := range a {
		v = v / 2
	}
	fmt.Println(a)
}

const (
	LINUX_REBOOT_MAGIC1      uintptr = 0xfee1dead
	LINUX_REBOOT_MAGIC2      uintptr = 672274793
	LINUX_REBOOT_CMD_RESTART uintptr = 0x1234567
)

// func reboot() {
// 	syscall.Syscall(syscall.SYS_REBOOT,
// 		LINUX_REBOOT_MAGIC1,
// 		LINUX_REBOOT_MAGIC2,
// 		LINUX_REBOOT_CMD_RESTART)
// }

func t4() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*3, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("found Match!")
	}

	re, _ := regexp.Compile(pat)

	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)

	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}

func t5() {
	// Here are some calculations with bigInts:
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip2 := big.NewInt(2)
	fmt.Println(ip.Mul(ip2, io))
	fmt.Println(ip.Mul(ip2, io).Add(ip, ip2))
	fmt.Println(ip)
	fmt.Println(ip.Mul(ip, ip2))
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)
	// Here are some calculations with bigInts:
	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Printf("Big Rat: %v\n", rq)
}

func main() {
	t5()
}
