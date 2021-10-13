package main

import (
	f "fmt"

	myUtil "example.com/mod_util"
)

func main() {
	f.Println("패키지 실습")
	a, b := 1, 2
	myUtil.Swap(&a, &b) //나만의 모듈?
	myToken := myUtil.Token{Name: "나의토큰", Total: 100000}
	f.Println(myToken)
	f.Println(a, b) // 2 1이라고 나와야함
}
