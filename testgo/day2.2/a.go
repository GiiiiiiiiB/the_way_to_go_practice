package main

import (
	. "fmt"
	"strings"
	// "unicode/utf8"
)

var a, b string = "asSASA ddd dsjkdsjs dk", "asSASA ddd dsjkdsjsこん dk"

func testRange() {
	a = "asSASA ddd 放假 dsjkdsjs dk"
	for _, val := range strings.Fields(a) {
		Printf("%s\n", val)
	}
	b = "GO1|The ABC of Go|25"
	str2 := strings.Split(b, "|")
	for _, val := range str2 {
		Printf("%s\n", val)
	}
	Printf("%v \n", str2)
}

func testReader() {
	reader1 := strings.NewReader(a)
	Println(reader1.ReadRune())
}

func main() {
	// Println(len(a))
	// Println(len(b))
	// Println(utf8.RuneCountInString(a))
	// Println(utf8.RuneCountInString(b))
	// Println(len([]rune(b)))

	str := "Hi, I'm Marc, 放假 Hi."

	Println(strings.Contains(str, "放假"))
	testRange()
	testReader()

}
