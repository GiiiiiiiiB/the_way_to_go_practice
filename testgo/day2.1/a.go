package main

import (
	"fmt"
	"math/rand"
	"unicode"
	// "time"
)

func t() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d /", a)
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d / ", r)
	}
	fmt.Println()
	// timens := int64(time.Now().Nanosecond())
	// rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.4f / ", 100*rand.Float32())
	}
	fmt.Println()
}

func dd() {
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U\n", ch, ch2, ch3) // UTF-8 code point
	fmt.Printf("%s - %s - %s\n", ch, ch2, ch3)
	fmt.Printf("%v", unicode.IsDigit(rune(ch)))
}

func aa() {
	a := "小明" //字节长度
	fmt.Println(len(a))

}

func main() {
	// var n int16 = 34
	// var m int32
	// // compiler error: cannot use n (type int16) as type int32 in assignment
	// //m = n
	// m = int32(n)

	// fmt.Printf("32 bit int is: %v\n", m)
	// fmt.Printf("16 bit int is: %v\n", n)
	aa()
}
