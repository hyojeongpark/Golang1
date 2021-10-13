package main

import "fmt"

func main() {
	var money int
	var menu string

	// 입력함수와 조건문 그리고 반복문을 이용해서 자판기를 만들어볼 생각.
	for {
		fmt.Printf("소지하고 있는 금액 입력 (10000 ~ 20000) : \n")

		_, err := fmt.Scanf("%d\n", &money)
		if err != nil {
			fmt.Println("잘못 입력했습니다.")
		} else {
			if money < 10000 || money > 20000 {
				//9,999원 이거나 20,001원일 경우 무한루프를 돌도록 설계
				continue
			} else {
				if money >= 10000 && money <= 19999 {
					break
				}
			}
		}
	}
	// 메뉴의 가격은 임믜로 하고 싶은 대로
	//for 문과 조건문을 이용해서 자판기를 만들어내는 main 함수를 만들어주세요.
	//무한 반복문 안에 menu 입력하는 함수 scanf를 이용해 메뉴를 입력받고
	// 메뉴의 금액에 따라 소지하고 있는 금액을 줄여 금액이 바닥 날때까지
	// 자판기를 돌려서 금액이 바닥나면 무한 반복문을 탈출하는 프로그램을 만들면 됩니다.
	for {
		// 메뉴 입력 받는 곳
		for {
			fmt.Print("마실 음료수: \n")

			_, err := fmt.Scanf("%s\n", &menu)

			if err != nil {
				fmt.Println("잘못 입력했습니다.")
			} else {
				break
			}
			// switch 문으로 case 별로 메뉴와 가격을 정해서 menu 값이 그 메뉴라면 소지금액에서 메뉴 가격을 -- 하면 됨.
			// 소지금액이 바닥나는 조건문은 if문을 통해 소지금액이 0보다 작으면 바깥 무한 반복문을 탈출하는 코드
		}
		switch menu {
		case "콜라":
			money = money - 200
			fmt.Printf("남은 소지금액은 %d입니다.", money)
		case "생수":
			money = money - 500
			fmt.Printf("남은 소지금액은 %d입니다.", money)
		case "게토레이":
			money = money - 1200
			fmt.Printf("남은 소지금액은 %d입니다.", money)
			//case 5개
		default:
			money = money - 1000
			fmt.Printf("남은 소지금액은 %d입니다.", money)
		}
		if money <= 0 {
			fmt.Println("금액이 부족합니다.")
			break
		}
	}
}
