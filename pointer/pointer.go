package main

import "fmt"

func main() {
	var num = 5
	var numPtr *int = &num //num=5의 주소인 numPtr
	p2 := &numPtr
	p3 := &p2
	fmt.Println("num값:", num, "num주소:", &num)
	fmt.Println("numPtr값:", numPtr, "numPtr주소:", &numPtr, "numPtr이 가리키는 값:", *numPtr)
	fmt.Println("p2값:", p2, "p2주소:", &p2, "p2가 가리키는 값:", *p2)
	fmt.Println("p3값:", p3, "p3주소:", &p3, "p3가 가리키는 값:", *p3)

}
