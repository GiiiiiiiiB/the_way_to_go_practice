package main

import (
	. "fmt" // Package implementing formatted I/O.
	"math"
	"os"
	"runtime"
)

type IZ = int
type IZ2 int

var ddd = "xiaoming"

const z = (1 << 100) >> 97

const (
	_  = iota             // 使用 _ 忽略不需要的 iota
	KB = 1 << (10 * iota) // 1 << (10*1)
	MB                    // 1 << (10*2)
	GB                    // 1 << (10*3)
	TB                    // 1 << (10*4)
	PB                    // 1 << (10*5)
	EB                    // 1 << (10*6)
	ZB                    // 1 << (10*7)
	YB                    // 1 << (10*8)
)

// const k = 1 / Ln2

func main() {
	Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
	var a IZ = 5
	var kk IZ = 5
	ak := kk
	var y = new(int)
	*y = 5
	Println(a)
	var b IZ2 = 10
	Println(b)
	Println(z)
	Println(KB, "KB")
	Println(MB, "MB")
	var d bool

	if a > 0 {
		d = true
	} else {
		d = false
	}
	Println(d)
	Println("=========================")
	x := 1
	Printf("this is %v\n", x)
	Printf("%v\n", x)
	Println("=========================")

	Println(runtime.GOOS)

	var goos string = runtime.GOOS

	Printf("The operating system is: %s\n", goos)
	Printf("%T\n", runtime.GOOS)
	Println(runtime.NumCPU())

	Println(os.Getenv("PATH"))
	Println(&a, y, &kk, &ak)
	Println(kk == *y)

	pi := math.Atan(1) * 4
	Println(pi)

}
