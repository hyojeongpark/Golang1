package main

import f "fmt"

func main() {
	//var num int
	//f.Println("1 ~ 10 까지의 정수 하나를 입력")
	//f.Scanln(&num)
	a := 10
	b := 20
	if show() && a > b {
		f.Println("Short-Circuit 호출")
	}
}
func show() bool {
	f.Println("Show 호출")
	return true
}

// if num >= 1 && num <= 5 {
// 	f.Println("숫자는 1보다 크고 5보다 작거나 같다")
// } else if num >= 6 && num <= 10 {
// 	f.Println("숫자는 6보다 크고 10보다 작거나 같다")
// } else {
// 	show()
// }
