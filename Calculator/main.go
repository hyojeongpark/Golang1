package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var op string

	fmt.Println("첫번째 숫자 입력하기")
	_, err1 := fmt.Scanf("%d\n", num1)

	if err1 != nil { //nil이 아님=
		fmt.Println("잘못입력했습니다")
	}
	fmt.Println("두번째 숫자 입력하기")
	_, err2 := fmt.Scanf("%d\n", num1)
	if err2 != nil {
		fmt.Println("잘못입력했습니다")
	}
	fmt.Println("연산자 입력하기(+, -, * ,/)")

	_, err3 := fmt.Scanf("%s\n", &op)
	if err3 != nil {
		fmt.Println("잘못입력했습니다")
	}

	// if op == "+" {
	// 	fmt.Printf("값은 %d", num1+num2)
	// } else if op == "-" {
	// 	fmt.Printf("값은 %d", num1-num2)
	// } else if op == "/" {
	// 	fmt.Printf("값은 %d", num1/num2)
	// } else if op == "*" {
	// 	fmt.Printf("값은 %d", num1*num2)
	// }

	switch op {
	case "+":
		fmt.Printf("값은 %d", num1+num2)
	case "-":
		fmt.Printf("값은 %d", num1-num2)
	case "/":
		fmt.Printf("값은 %d", num1/num2)
	case "*":
		fmt.Printf("값은 %d", num1*num2)
	default:
		fmt.Printf("연산자가 아닙니다")
	}
}
