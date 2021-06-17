package main

import "fmt"

func main() {
	var a = 123
	var b = 456
	var c = 123.45678
	var d = 101
	var e float32 = 324.15
	var str string = "ABC"

	var i int
	fmt.Scanln(&i)
	fmt.Printf("i 번지의 주소는.. %p\n", &i)
	fmt.Println("i의 실제값은..", i)

	fmt.Printf("%5d, %10d\n", a, b)
	fmt.Printf("%05d, %10d\n", a, b)
	fmt.Printf("%-5d, %-5d\n", a, b)
	fmt.Printf("%08.1f", c)
	fmt.Printf("\"안녕하세요\t다른내용입니다\"	%s,%f,%x", str, e, d)

}
