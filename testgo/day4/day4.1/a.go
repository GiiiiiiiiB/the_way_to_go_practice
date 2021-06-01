package main

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"strings"
	"time"
	"unicode"
)

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {

	fmt.Println("in b")
	defer un(trace("b"))
	a()
}

func func1(s string) (n int, err error) {
	a := 3
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
		log.Println(a)
	}()
	a = 4
	return 8, io.EOF
}

func func2() {
	i := 0
	defer fmt.Println(i)
	defer func() {
		fmt.Println(i)
	}()
	i = 1
	return
}

func f(a int) (k int, res int) {
	k = a
	if a <= 1 {
		res = 1
	} else {
		_, one := f(a - 1)
		_, two := f(a - 2)
		res = one + two
	}
	return
}

func dg(a int) {
	if a > 0 {
		fmt.Println(a)
		dg(a - 1)
	} else {
		return
	}
}

func chen(a int) int {
	if a == 1 {
		return a
	} else {
		return a * (chen(a - 1))
	}
}

func niming() {
	a := func() { fmt.Println("xxxxxx") }
	a()
	a()
}

func main() {
	// b()
	start := time.Now()
	func1("go")
	func2()
	a, b := f(10)
	fmt.Println(a, b)
	dg(10)
	fmt.Println(chen(15))
	c := strings.IndexFunc("nihao china", unicode.IsSpace)
	fmt.Println(c)
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	niming()
	where()
	log.Print("11111111111")
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)
}
