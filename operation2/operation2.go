package main

import "fmt"

func main() {
	a := "홍춘삼"
	b := "홍길동"
	if a > b {
		fmt.Printf("a인 %s가 더 크다\n", a)
	} else {
		fmt.Println("b인 %s가 더 크다\n", b)
	}
	c := 1
	//var d int
	c++
	d := c
	//컴파일에러-> d:=c++ 또는 d:=++c
	//다른변수에 넣을 수 없음(자바,c++은 됨)
	fmt.Println(c, d)
}
