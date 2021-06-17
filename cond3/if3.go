package main

import "fmt"

func main() {
	var score = 87 //여기엔 왜 int 쓰면 오류나지? type=var이기 때문
	fmt.Scan(&score)
	switch {
	case score > 90:
		fmt.Println("1 이하")
		fallthrough
	case score > 80:
		fmt.Println("2 이하")
		fallthrough
	case score > 70:
		fmt.Println("3 이하")
		fallthrough
	default:
		fmt.Println("default 도달")
	}
}
