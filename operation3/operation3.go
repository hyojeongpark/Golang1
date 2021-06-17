package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scan(&a, &b)
	// fmt.Printf("a + b = %d\n", a+b)
	// fmt.Printf("a - b = %d\n", a-b)
	// fmt.Printf("a * b = %d\n", a*b)
	// fmt.Printf("a / b = %d\n", a/b)
	// fmt.Printf("a %% b = %d\n", a%b)
	//fmt.Printf("%d %% %d = %d\n", a, b, a%b)
	//fmt.Scanf("%d\n%d", &a, &b) //엔터하는경우
	//fmt.Scanf("%d %d", &a, &b) //한 줄에 10 20 과 같이 입력할 때
	fmt.Printf("a+b=%d\n", a+b) //8번쨰줄 (&a, &b를 쓰는 경우)
}
